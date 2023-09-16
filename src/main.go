package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init(){
	err := godotenv.Load(".env")
	if err != nil {
			log.Fatal("Error loading .env file")
	}
}

func main() {
	token := getTwitchToken()
	fmt.Printf("Token Acquired: %s\n", token)

	game := getGameByName(token, "league of legends")
	fmt.Printf("Game ID: %s\n", game.ID)

	clips := getClipsByGame(token, game.ID, 1, "2023-09-13", "2023-09-15")
	filteredClips := filterClips(clips, "en", 10)

	directory, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	path := directory + "/downloads/"

	os.RemoveAll(path)
	os.Mkdir(path, 0700)

	for index, clip := range filteredClips {
		fmt.Printf("Clip %d: %s\n", index, clip.URL)

	 	filename := fmt.Sprintf("Clip%d", index)	
		downloadFileFromURL(path, filename, strings.Replace(clip.ThumbnailURL, "-preview-480x272.jpg", ".mp4", 1))
	}

	filterClipsPath := path + "/filtered"
	os.RemoveAll(filterClipsPath)
	os.Mkdir(filterClipsPath, 0700)

	for index, clip := range filteredClips {
		filename := fmt.Sprintf("Clip%d", index)	
		addStreamerToClip(path, filename, clip)	
	}
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