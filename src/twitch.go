package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

func getToken() string {
	
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

	// println("Access Token:", tokenResponse.AccessToken)
	// println("Expires In:", tokenResponse.ExpiresIn)
	// println("Token Type:", tokenResponse.TokenType)

	return tokenResponse.AccessToken
}