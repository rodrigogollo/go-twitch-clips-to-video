package main

import (
	"fmt"
	"log"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func addStreamerToClip(path, filename string, clip Clip){
	fullpath := path + filename + ".mp4"
	outputPath := path + "filtered/" + filename + ".mp4"

	input := ffmpeg.Input(fullpath)

	video := input.Filter("drawtext", ffmpeg.Args{fmt.Sprintf("text:%s", clip.BroadcasterName), 
	"enable=if(lt(t,0.3),gte(t,0.0),if(lt(t,1.3),gte(t,0.3),if(lt(t,3),gte(t,1.3),if(lt(t,4),gte(t,3),0))))",
	"x=15",
	"y=15", 
	"fontsize=50",
	"fontcolor=E8E2F3", 
	"box=1", 
	"boxcolor=6441A4", 
	"borderw=4",
	"bordercolor=1B112C", 
	"boxborderw=15",
	"fontfile=Roboto.ttf"}).
	Filter("fade", ffmpeg.Args{"t=in", "st=0", "d=1"}).
	Filter("fade", ffmpeg.Args{"t=out", fmt.Sprintf("st=%.2f", clip.Duration), "d=3"})

	audio := input.Audio() 
	streams := []*ffmpeg.Stream{video, audio}
  err := ffmpeg.Output(streams, outputPath, ffmpeg.KwArgs{"preset": "veryfast", "crf": "16", "vcodec": "libx264", "acodec": "aac", "ar": "48000"}).Run()

	if err != nil {
    log.Fatal(err) 
  }
}