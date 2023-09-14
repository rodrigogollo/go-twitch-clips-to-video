package main

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
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

type Clip struct {
	ID              string  `json:"id"`
	URL             string  `json:"url"`
	EmbedURL        string  `json:"embed_url"`
	BroadcasterID   string  `json:"broadcaster_id"`
	BroadcasterName string  `json:"broadcaster_name"`
	CreatorID       string  `json:"creator_id"`
	CreatorName     string  `json:"creator_name"`
	VideoID         string  `json:"video_id"`
	GameID          string  `json:"game_id"`
	Language        string  `json:"language"`
	Title           string  `json:"title"`
	ViewCount       int     `json:"view_count"`
	CreatedAt       string  `json:"created_at"`
	ThumbnailURL    string  `json:"thumbnail_url"`
	Duration        float64 `json:"duration"`
	VodOffset       int     `json:"vod_offset"`
	IsFeatured      bool    `json:"is_featured"`
}

type ClipsResponse struct {
	Data []Clip `json:"data"`
}