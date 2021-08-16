package main


// Import Ebiten v2, and pearl
import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/eboatwright/pearl"
)


// Constants for window size, screen scale, and title
// Screen scale is how many screen pixels, are in game pixels (for pixel art)
const (
	WINDOW_WIDTH  = 960
	WINDOW_HEIGHT = 600
	SCREEN_SCALE  = 3
	WINDOW_TITLE  = "Pearl Engine"
)

// This isn't a constant, because color.RGBA isn't a const initializer. But, I'll use it like a constant anyway. 😎
var (
	BACKGROUND_COLOR = pearl.RGBA(127, 127, 127, 255)
)


// Transform component, has a pearl Vector2 as a position
type Transform struct {
	position pearl.Vector2
}
// return the name of the component
func (t *Transform) ID() string { return "transform" }

type Image struct {
	image *ebiten.Image
}
func (t *Image) ID() string { return "image" }

type Player struct {
	velocity pearl.Vector2
	speed    float64
	friction float64
}
func (p *Player) ID() string { return "player" }


// This is a system
type ImageSystem struct {}

// Use this for updating entities
func (is *ImageSystem) Update(entity *pearl.Entity, scene *pearl.Scene) {}

// Use this for drawing entities
func (is *ImageSystem) Draw(entity *pearl.Entity, scene *pearl.Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	// Get the components
    t := entity.GetComponent("transform").(*Transform)
    i := entity.GetComponent("image").(*Image)

	// This just sets the options.GeoM to be the position specified
    pearl.SetDrawPosition(options, t.position)
	// Draw the image onto the ebiten screen
    screen.DrawImage(i.image, options)
}

// Return the components you need for this system
func (is *ImageSystem) GetRequirements() []string {
    return []string { "transform", "image" }
}

type PlayerSystem struct {}

// This system utilizes the Update loop
func (ps *PlayerSystem) Update(entity *pearl.Entity, scene *pearl.Scene) {
	p := entity.GetComponent("player").(*Player)
	t := entity.GetComponent("transform").(*Transform)

	// Creates a Vector2 with X and Y being input directions
	input := pearl.Vector2 {
		float64(pearl.GetInputAxis([]ebiten.Key { ebiten.KeyA }, []ebiten.Key { ebiten.KeyD })),
		float64(pearl.GetInputAxis([]ebiten.Key { ebiten.KeyW }, []ebiten.Key { ebiten.KeyS })),
	}
	// Normalize input (this keeps the player for going alot faster if going diagonal)
	input.Normalize()

	// Add the velocity multiplied by players speed
	p.velocity.Add(pearl.Vector2Multiply(input, p.speed))
	// Multiply friction to velocity
	p.velocity.Multiply(p.friction)
	// And finally add velocity (floored) to position (flooring velocity helps fix jittering)
	t.position.Add(pearl.Vector2Floor(p.velocity))
}

func (ps *PlayerSystem) Draw(entity *pearl.Entity, scene *pearl.Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions) {}

func (ps *PlayerSystem) GetRequirements() []string {
    return []string { "player", "transform" }
}


// Function called when pearl starts!
func onStart() {
	// Create a scene and load it
	gameScene := &pearl.Scene{ ID: "game" }
	pearl.LoadScene(gameScene)

	// Create an entity
	player := &pearl.Entity{ ID: "test" }
	// Add components
	player.AddComponents([]pearl.IComponent {
		&Transform {
			position: pearl.Vector2 { 50, 50 },
		},
		&Player {
			speed:    0.8,
			friction: 0.8,
		},
		&Image {
			image: pearl.LoadImage("data/img/wizard.png"),
		},
	})

	// Add entities to game scene
	gameScene.AddEntities([]*pearl.Entity {
		player,
	})

	// Add an Image System to game scene
	gameScene.AddSystems([]pearl.ISystem {
		&PlayerSystem {  },
		&ImageSystem {  },
	})
}


func main() {
	// Start pearl!
	// The last parameter, is the function that will be called when pearl is ready
	pearl.Start(WINDOW_WIDTH, WINDOW_HEIGHT, SCREEN_SCALE, WINDOW_TITLE, BACKGROUND_COLOR, onStart)
}