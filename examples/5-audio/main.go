package main


import (
	"github.com/eboatwright/pearl"
)


func main() {
	pearl.Start(960, 600, 3, "Audio Test", pearl.RGBA(0, 0, 0, 255), func() {
		scene := &pearl.Scene { ID: "scene" }
		pearl.LoadScene(scene)
		pearl.PlaySoundFromFile("test.wav")
	})
}