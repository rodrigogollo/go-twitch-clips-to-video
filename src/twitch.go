package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func getTwitchToken() string {
	
	type TokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}

	CLIENT_ID := os.Getenv("CLIENT_ID")
	CLIENT_SECRET := os.Getenv("CLIENT_SECRET")

	data := url.Values{}
	data.Set("client_id", CLIENT_ID)
	data.Set("client_secret", CLIENT_SECRET)
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://id.twitch.tv/oauth2/token", bytes.NewBuffer([]byte(data.Encode())))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		panic(err)
	}

  client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)

	if err != nil {
		panic(err)
	}

	println("Access Token:", tokenResponse.AccessToken)
	println("Expires In:", tokenResponse.ExpiresIn)
	println("Token Type:", tokenResponse.TokenType)

	return tokenResponse.AccessToken
}

func getClipsByGame(token, name string) {

	// gameData := getGameByName(token, name)


}

type GamesResponse struct {
	Data []Game `json:"data"`
}

type Game struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BoxArtUrl string `json:"box_art_url"`
	IGDB_ID   string `json:"igdb_id"`
}

func getGameByName(token, name string) Game {

	CLIENT_ID := os.Getenv("CLIENT_ID")

	baseURL, _ := url.Parse("https://api.twitch.tv/helix/games")
	params := url.Values{}
	params.Add("name", "cs2")
	baseURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	req.Header.Set("Client-Id", CLIENT_ID)

	if err != nil {
		panic(err)
	}

  client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)

	var gamesResponse GamesResponse
	err = json.Unmarshal([]byte(body), &gamesResponse)

	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))

	games := gamesResponse.Data
	firstGame := games[0]

	return firstGame
}