package main

type MediaPlayer interface {
	play(mediaType string, fileName string)
}

type Mp4Player struct {
	MediaPlayer
}

type AviPlayer struct {
	MediaPlayer
}

type MediaAdapter struct {
	MediaPlayer
}

type AudioPlayer struct {
	MediaPlayer
}

func (p *Mp4Player) play(fileName string) {
	println("play mp4")
}

func (p *AviPlayer) play(fileName string) {
	println("play avi")
}

func (a *MediaAdapter) play(mediaType string, fileName string) {
	switch mediaType {
	case "avi":
		var player AviPlayer
		player.play(fileName)
	case "mp4":
		var player Mp4Player
		player.play(fileName)
	default:
		panic("mediaType error")
	}
}

func (a *AudioPlayer) play(mediaType string, fileName string) {
	switch mediaType {
	case "mp3":
		println("play mp3")
	default:
		var adapter MediaAdapter
		adapter.play(mediaType, fileName)
	}
}
