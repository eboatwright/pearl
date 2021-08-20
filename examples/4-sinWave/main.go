package main


import (
	"github.com/eboatwright/pearl"
	bicas "github.com/eboatwright/pearl/builtInComponentsAndSystems"
)


func main() {
	pearl.Start(960, 600, 1, "Sin Wave", pearl.RGBA(0, 0, 0, 255), func() {
		scene := &pearl.Scene { ID: "scene" }
		pearl.LoadScene(scene)

		rect := &pearl.Entity { ID: "rect" }
		rect.AddComponents([]pearl.IComponent {
			&bicas.Transform {
				Position: pearl.Vector2 { 100, 100 },
				Scale:    pearl.Vector2 { 100, 100 },
			},
			&bicas.RectangleRenderer {
				Color: pearl.RGBA(255, 255, 255, 255),
			},
			&bicas.SinWave {
				Speed:    1,
				Distance: 50,
			},
		})

		scene.AddEntities([]*pearl.Entity {
			rect,
		})

		scene.AddSystems([]pearl.ISystem {
			&bicas.RectangleRendererSystem {  },
		})
	})
}