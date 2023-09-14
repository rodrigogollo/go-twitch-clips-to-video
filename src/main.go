package main

import (
	"fmt"
	"log"

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

	game := getGameByName(token, "cs2")
	fmt.Printf("Game ID: %s\n", game.ID)


	// makevideo("broadcast", "jerma985", "day", 10)
}

func makevideo(category, name, period string, size int) {
	switch category {
	case "game":
		fmt.Println("game")
	case "broadcast": 
		fmt.Println("broadcast")
	}

	fmt.Println(name, period, size)

}