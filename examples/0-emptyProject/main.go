package main


// Import Ebiten and Pearl (Ebiten not used in empty example)
import (
	// "github.com/hajimehoshi/ebiten/v2"
	"github.com/eboatwright/pearl"
)


// Define window size, scale and title
const (
	WINDOW_WIDTH  = 960
	WINDOW_HEIGHT = 600
	SCREEN_SCALE  = 3
	WINDOW_TITLE  = "empty pearl"
)

// This isn't const because pearl.RGBA isn't a const initializer
var (
	BACKGROUND_COLOR = pearl.RGBA(135, 206, 235, 255)
)


// Called when Pearl is ready
func onStart() {
	// Create and load scene
	gameScene := &pearl.Scene { ID: "game" }
	pearl.LoadScene(gameScene)
}


// Entry point
func main() {
	// Start Pearl with predefined constants
	// Also, pass in function to be called on start
	pearl.Start(WINDOW_WIDTH, WINDOW_HEIGHT, SCREEN_SCALE, WINDOW_TITLE, BACKGROUND_COLOR, onStart)
}