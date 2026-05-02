package main

import "fmt"

type Player interface {
	Play() string
}

type Audio struct {
	title string
}

type Video struct {
	title      string
	resolution string
}

func (audio Audio) Play() string {
	return fmt.Sprintf("audio: %s is playing", audio.title)
}

func (video Video) Play() string {
	return fmt.Sprintf("video of resolution %s: %s is playing", video.resolution, video.title)
}

func main() {
	playlist := []Player{
		Audio{title: "Go Podcast Ep 1"},
		Video{title: "Go Interfaces Tutorial", resolution: "1080p"},
		Audio{title: "Lofi Coding Beats"},
	}
	for _, media := range playlist {
		fmt.Println(media.Play())
	}
}
