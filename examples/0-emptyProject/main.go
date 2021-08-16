package main


import (
	// "github.com/hajimehoshi/ebiten/v2"
	"github.com/eboatwright/pearl"
)


const (
	WINDOW_WIDTH  = 960
	WINDOW_HEIGHT = 600
	SCREEN_SCALE  = 3
	WINDOW_TITLE  = "empty pearl"
)

var (
	BACKGROUND_COLOR = pearl.RGBA(135, 206, 235, 255)
)


func onStart() {
	gameScene := &pearl.Scene { ID: "game" }
	pearl.LoadScene(gameScene)
}


func main() {
	pearl.Start(WINDOW_WIDTH, WINDOW_HEIGHT, SCREEN_SCALE, WINDOW_TITLE, BACKGROUND_COLOR, onStart)
}