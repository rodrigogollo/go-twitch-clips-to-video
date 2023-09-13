package main

import (
	"fmt"
)

func main() {
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