package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load(".env")
	if err != nil {
			log.Fatal("Error loading .env file")
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter the game name: ")
	gameName, _ := reader.ReadString('\n')
	gameName = strings.TrimSpace(gameName)

	fmt.Print("Enter the size of clips: ")
	clipSizeInput, _ := reader.ReadString('\n')
	clipSizeInput = strings.TrimSpace(clipSizeInput)

	clipSize, err := strconv.Atoi(clipSizeInput)
	if err != nil {
		fmt.Println("Invalid input for clip size. Please enter a valid number.")
		return
	}

	fmt.Print("Enter the number of days before today to search for clips: ")

	daysInput, _ := reader.ReadString('\n')
	daysInput = strings.TrimSpace(daysInput)

	days, err := strconv.Atoi(daysInput)
	if err != nil {
		fmt.Println("Invalid input for clip size. Please enter a valid number.")
		return
	}

	fmt.Printf("You entered the game name: %s\n", gameName)
	fmt.Printf("You entered the clip size: %d\n", clipSize)
	fmt.Printf("You entered the days before: %d\n", days)
	fmt.Printf("Searching %d Twitch Clips for game %s in the last %d days.\n", clipSize, gameName, days)

	getClipsAndDownload(gameName, clipSize, days)
	
	// wanted := []int{1, 4, 5, 8, 10}
	// addFilterToWantedClips(wanted)
	// mergeClips("dayz", wanted)
}

func getClipsAndDownload(gamename string, size, days int) {
	token := getTwitchToken()
	fmt.Printf("Token Acquired: %s\n", token)

	game := getGameByName(token, gamename)

	dateEnd := time.Now().UTC()
	dateStart := dateEnd.AddDate(0, 0, -days)

	clips := getClipsByGame(token, game.ID, size, dateStart.Format("2006-01-02"), dateEnd.Format("2006-01-02"))
	filteredClips := filterClips(clips, "en", 25)
	
	jsonInfo, _ := json.Marshal(filteredClips)
	os.WriteFile("output.json", jsonInfo, 0644)

	directory, _ := os.Getwd()
	path := directory + "/downloads/"

	os.RemoveAll(path)
	os.Mkdir(path, 0700)

	var wg sync.WaitGroup

	wg.Add(len(filteredClips))
	for index, clip := range filteredClips {
	go func(i int, clip Clip){
		defer wg.Done()
			fmt.Printf("Clip %d: %s\n", i, clip.URL)
	
		 	filename := fmt.Sprintf("Clip%d", i)	
			downloadFileFromURL(path, filename, strings.Replace(clip.ThumbnailURL, "-preview-480x272.jpg", ".mp4", 1))
		}(index, clip)
	}
	wg.Wait()
}

func addFilterToWantedClips(wanted []int) {
	file, err := os.ReadFile("output.json")

	if err != nil {
		fmt.Println(err)
	}

	filteredClips := []Clip{}
	err = json.Unmarshal([]byte(file), &filteredClips)

	if err != nil {
		fmt.Println(err)
	}

	var slice []*Clip
	for _, wantedClip := range wanted {
		slice = append(slice, &filteredClips[wantedClip])
	}

	directory, _ := os.Getwd()
	path := directory + "/downloads/"

	filterClipsPath := path + "/filtered"
	os.RemoveAll(filterClipsPath)
	os.Mkdir(filterClipsPath, 0700)

	var wg sync.WaitGroup
	wg.Add(len(slice))
	for index, clip := range slice {
	go func(i int, clip *Clip){
		defer wg.Done()
			filename := fmt.Sprintf("Clip%d", wanted[i])	
			addStreamerToClip(path, filename, clip)	
		}(index, clip)
	}
	wg.Wait()	
}

func downloadFileFromURL(path, filename, url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	filepath := fmt.Sprintf(`%s/%s`, path, filename)
	out, err := os.Create(filepath + ".mp4")

	if err != nil {
		panic(err)
	}

	size, err := io.Copy(out, resp.Body)
	
	if err != nil {
		panic(err)
	}

	defer out.Close()

	fmt.Printf("Downloaded a file %s with size %d\n", filepath, size)
}