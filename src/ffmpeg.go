package main

import (
	"fmt"
	"log"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func addStreamerToClip(path, filename string, clip *Clip){
	fullpath := path + filename + ".mp4"
	outputPath := path + "filtered/" + filename + ".mp4"

	input := ffmpeg.Input(fullpath)

	video := input.Filter("drawtext", ffmpeg.Args{fmt.Sprintf("text:%s", clip.BroadcasterName), 
	"enable=if(lt(t,3),gte(t,0.0),if(lt(t,1.3),gte(t,3),if(lt(t,3),gte(t,1.3),if(lt(t,4),gte(t,3),0))))",
	"x=15",
	"y=15", 
	"fontsize=56",
	"fontcolor=E8E2F3", 
	"box=1", 
	"boxcolor=6441A4", 
	"borderw=4",
	"bordercolor=1B112C", 
	"boxborderw=15",
	"fontfile=Roboto.ttf"}).
	Filter("fade", ffmpeg.Args{"t=in", "st=0", "d=1"}).
	Filter("fade", ffmpeg.Args{"t=out", fmt.Sprintf("st=%.2f", clip.Duration - 2), "d=2"}).
	Filter("scale", ffmpeg.Args{"w=1920", "h=1080"})

	audio := input.Audio() 
	streams := []*ffmpeg.Stream{video, audio}
  err := ffmpeg.Output(streams, outputPath, ffmpeg.KwArgs{"preset": "veryfast", "crf": "16", "vcodec": "libx264", "acodec": "aac", "ar": "48000"}).Run()

	if err != nil {
    log.Fatal(err) 
  }
}

func mergeClips(channel string, wanted []int) {

	directory, _ := os.Getwd()

	streams := []*ffmpeg.Stream{}

	channelPath := directory + "/content/" + channel
	intro := channelPath + "/shortintro.mp4"
	outro := channelPath + "/outro.mp4"

	streams = append(streams, ffmpeg.Input(intro).Video())
	streams = append(streams, ffmpeg.Input(intro).Audio())

	for _, inputClip := range wanted {
		inputPath := directory + "/downloads/filtered/" + fmt.Sprintf("Clip%d", inputClip) + ".mp4"
		fmt.Printf("stream %d - %s\n", inputClip, inputPath)
		
		input := ffmpeg.Input(inputPath)

		streams = append(streams, input.Video())
		streams = append(streams, input.Audio())
	}

	streams = append(streams, ffmpeg.Input(outro).Video())
	streams = append(streams, ffmpeg.Input(outro).Audio())

  err := ffmpeg.Concat(streams, ffmpeg.KwArgs{"v": 1, "a": 1}).
	Output("output.mp4", ffmpeg.KwArgs{"preset": "veryfast", "crf": "16", "vcodec": "libx264", "acodec": "aac", "ar": "48000", "r": "30"}).
	OverWriteOutput().
	ErrorToStdOut().
	Run()
	
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}