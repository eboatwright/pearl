package main


import (
	"github.com/eboatwright/pearl"
	bicas "github.com/eboatwright/pearl/builtInComponentsAndSystems"
)


func main() {
	pearl.Start(960, 600, 3, "Rectangle Rendering", pearl.RGBA(0, 0, 0, 255), func() {
		gameScene := &pearl.Scene { ID: "game" }
		pearl.LoadScene(gameScene)

		rect := &pearl.Entity { ID: "rect" }
		rect.AddComponents([]pearl.IComponent {
			&bicas.Transform {
				Position: pearl.Vector2 { 5, 5 },
				Scale:    pearl.Vector2 { 10, 10 },
			},
			&bicas.RectangleRenderer {
				Color: pearl.RGBA(255, 255, 255, 255),
			},
		})

		gameScene.AddEntities([]*pearl.Entity {
			rect,
		})

		gameScene.AddSystems([]pearl.ISystem {
			&bicas.RectangleRendererSystem {  },
		})
	})
}