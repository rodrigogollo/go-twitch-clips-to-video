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
	token := getToken()
	fmt.Printf("Token Acquired: %s", token)
	makevideo("broadcast", "jerma985", "day", 10)
	// resp, err := http.Get("https://dummyjson.com/users")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf(string(body))

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