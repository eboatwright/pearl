package pearl


import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/audio"

	"image/color"
	"fmt"
	"time"
)


type function func()

type game struct {
	windowWidth     int
	windowHeight    int
	screenScale     int
	windowTitle     string
	backgroundColor color.Color
	currentScene    *Scene
	showFPS         bool
	timeStart       time.Time
	audioContext    *audio.Context
}

func (g *game) Update() error {
	updateInput()
	g.currentScene.Update()
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(g.backgroundColor)

	g.currentScene.Draw(screen)

	if g.showFPS {
		ebitenutil.DebugPrint(screen, "FPS: " + fmt.Sprint(int32(ebiten.CurrentFPS())))
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.windowWidth / g.screenScale, g.windowHeight / g.screenScale
}


// Makes scene specified the current scene
func LoadScene(scene *Scene) {
	g.currentScene = scene
}

// Call this to toggle the FPS overlay
func ToggleFPS() {
	g.showFPS = !g.showFPS
}

// Returns screen size as Vector2
func GetScreenSize() Vector2 {
	return Vector2 { float64(g.windowWidth / g.screenScale), float64(g.windowHeight / g.screenScale) }
}

// Returns the time since the program started
func TimeSinceStart() float64 {
	return float64(time.Since(g.timeStart).Milliseconds())
}

var g *game
// Call this to start Pearl! :D
func Start(windowWidth, windowHeight, screenScale int, windowTitle string, backgroundColor color.Color, onStart function) {
	g = &game {
		windowWidth:     windowWidth,
		windowHeight:    windowHeight,
		screenScale:     screenScale,
		windowTitle:     windowTitle,
		backgroundColor: backgroundColor,
		timeStart:       time.Now(),
		audioContext:    audio.NewContext(48000),
	}

	ebiten.SetWindowSize(g.windowWidth, g.windowHeight)
	ebiten.SetWindowTitle(g.windowTitle)

	onStart()

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}