package main


import (
	// "github.com/hajimehoshi/ebiten/v2"
	"github.com/eboatwright/pearl"
	bicas "github.com/eboatwright/pearl/builtInComponentsAndSystems"
)


const (
	WINDOW_WIDTH  = 960
	WINDOW_HEIGHT = 600
	SCREEN_SCALE  = 3
	WINDOW_TITLE  = "Pearl Platformer"
)

var (
	BACKGROUND_COLOR = pearl.RGBA(135, 206, 235, 255)
)


func onStart() {
	gameScene := &pearl.Scene { ID: "game" }
	pearl.LoadScene(gameScene)

	player := &pearl.Entity { ID: "player" }
	player.AddComponents([]pearl.IComponent {
		&bicas.Transform {
			Position: pearl.Vector2 { 8, 8 },
			Scale:    pearl.VONE,
		},
		&bicas.Image {
			Image: pearl.LoadImage("data/img/wizard.png"),
			Size:  pearl.Vector2 { 16, 16 },
		},
		&bicas.Rigidbody {
			Friction:   pearl.Vector2 { 0.8, 1 },
			Gravity:    pearl.Vector2 { 0, 0.4 },
			CoyoteTime: 3,
		},
	})

	gameScene.AddEntities([]*pearl.Entity {
		player,
	})

	gameScene.AddSystems([]pearl.ISystem {
		&bicas.NonCollisionRigidbodySystem {  },

		&bicas.LevelRenderer {  },
		&bicas.ImageRenderer {  },
	})
}


func main() {
	pearl.Start(WINDOW_WIDTH, WINDOW_HEIGHT, SCREEN_SCALE, WINDOW_TITLE, BACKGROUND_COLOR, onStart)
}