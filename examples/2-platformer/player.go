package main


import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/eboatwright/pearl"
	bicas "github.com/eboatwright/pearl/builtInComponentsAndSystems"
)


type Player struct {
	MoveSpeed  float64
	JumpHeight float64
}
func (p *Player) ID() string { return "player" }


type PlayerSystem struct {}

func (ps *PlayerSystem) Update(entity *pearl.Entity, scene *pearl.Scene) {
	rb := entity.GetComponent("rigidbody").(*bicas.Rigidbody)
	p := entity.GetComponent("player").(*Player)

	xInput := pearl.GetInputAxis(
		[]ebiten.Key { ebiten.KeyA, ebiten.KeyLeft },
		[]ebiten.Key { ebiten.KeyD, ebiten.KeyRight },
	)

	rb.Velocity.X += float64(xInput) * p.MoveSpeed

	if pearl.KeyJustPressed(ebiten.KeyW) || pearl.KeyJustPressed(ebiten.KeyUp) {
		rb.Grounded = 0
		rb.Velocity.Y = p.JumpHeight
	}
}

func (ps *PlayerSystem) Draw(entity *pearl.Entity, scene *pearl.Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions) {}

func (ps *PlayerSystem) GetRequirements() []string {
	return []string { "rigidbody", "player" }
}