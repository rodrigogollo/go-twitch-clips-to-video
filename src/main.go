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

	game := getGameByName(token, "dayz")
	fmt.Printf("Game ID: %s\n", game.ID)

	clips := getClipsByGame(token, game.ID, 1, "2023-08-10", "2023-09-14")
	filteredClips := filterClips(clips, "en", 10)

	for index, clip := range filteredClips {
		fmt.Printf("Clip %d: %s\n", index, clip.URL)

		directory, err := os.Getwd()
		if err != nil {
      fmt.Println(err) //print the error if obtained
   }
		downloadFileFromURL(fmt.Sprintf(`%s/%s`, directory, "downloads"), fmt.Sprintf("Clip%d.mp4", index), strings.Replace(clip.ThumbnailURL, "-preview-480x272.jpg", ".mp4", 1))
}
	


	// makevideo("broadcast", "jerma985", "day", 10)
}


func downloadFileFromURL(path, name, url string) error{
	os.RemoveAll(path)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	filepath := fmt.Sprintf(`%s/%s`, path, name)
	out, err := os.Create(filepath)

	if err != nil {
		return err
	}

	size, err := io.Copy(out, resp.Body)

	defer out.Close()

	fmt.Printf("Downloaded a file %s with size %d\n", filepath, size)
	return err
}