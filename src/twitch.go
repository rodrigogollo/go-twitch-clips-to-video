package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func getTwitchToken() string {
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

func getClipsByGame(token string, gameId string, first int, startedAt string, endedAt string) []Clip {
	CLIENT_ID := os.Getenv("CLIENT_ID")
	baseURL, _ := url.Parse("https://api.twitch.tv/helix/clips")
	
	startedAtDate, err := time.Parse("2006-01-02", startedAt)
	if err != nil {
		panic(err)
	}
	
	endedAtDate, err := time.Parse("2006-01-02", endedAt)
	if err != nil {
			panic(err)
	}

	startDate := startedAtDate.Format("2006-01-02T15:04:05Z")
	endDate := endedAtDate.Format("2006-01-02T15:04:05Z")

	params := url.Values{}
	params.Set("game_id", gameId)
	params.Set("first", strconv.Itoa(first))
	params.Set("started_at", startDate)
	params.Set("ended_at", endDate)

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

	// fmt.Println(string(body))

	var clipsResponse ClipsResponse
	err = json.Unmarshal([]byte(body), &clipsResponse)

	if err != nil {
		panic(err)
	}

	clips := clipsResponse.Data
	return clips

}