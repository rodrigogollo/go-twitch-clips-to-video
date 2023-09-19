# Twitch Clips to Video (Refactor in Golang)

## This project is a tool to automatically create compilation videos from Twitch clips for a specific game. It consists of a Golang backend that queries the Twitch API to find recent clips for a game, and the FFMPEG go package for video editing scripts that filters the clips and stitches together clips into a full video.

#### This project is a refactor of my other project build in NodeJS. I did the refactor because the other project was built intending to use the Youtube API and Remotion Video Editor. And it didn't work that well as I first thought it would. So I decided to refactor to learn more about the differences between Golang and JavaScript. And learn more about Golang development in general.

### Features
- Fetch recent clips for a game from Twitch API.
- Download clips locally.
- Add video filter with the broadcaster name in it.
- Compile and stitch clips into a full video.
- Output video in mp4 format.

### Technologies:
- Golang (fmt, io, net/http, os, sync, time, encoding/json)
- FFMPEG-GO package (github.com/u2takey/ffmpeg-go)
- Twitch API (Token Auth and Helix Games)

Run the commmand: 
```bash 
go run ./src/
```

Input the Parameters

```bash 
Game name: Fortnite
Clip size: 20 
Days before: 7
```

Then the clip will start to download locally:

![Alt text](image.png)