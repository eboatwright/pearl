package main


import (
	"github.com/eboatwright/pearl"
)


const (
	WINDOW_WIDTH  = 960
	WINDOW_HEIGHT = 600
	SCREEN_SCALE  = 3
	WINDOW_TITLE  = "Testing Pearl"
)

var (
	BACKGROUND_COLOR = pearl.RGBA(255, 255, 255, 255)
)


func onStart() {
	println("on start!")
}


func main() {
	pearl.Start(WINDOW_WIDTH, WINDOW_HEIGHT, SCREEN_SCALE, WINDOW_TITLE, BACKGROUND_COLOR, onStart)
}