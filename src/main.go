package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	// gamename := "dayz"
	// getClipsAndDownload(gamename, 50, 5)

	wanted := []int{5, 4, 11,7, 10}
	addFilterToWantedClips(wanted)

	mergeClips("dayz", wanted)
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
	wg.Add(3)
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