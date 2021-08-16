package builtInComponentsAndSystems


import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/eboatwright/pearl"
)


type Rigidbody struct {
	Velocity   pearl.Vector2
	Friction   pearl.Vector2
	Gravity    pearl.Vector2
	Grounded   float32
	CoyoteTime float32
}
func (rb *Rigidbody) ID() string { return "rigidbody" }


type NonCollisionRigidbodySystem struct {}

func (rbs *NonCollisionRigidbodySystem) Update(entity *pearl.Entity, scene *pearl.Scene) {
	t := entity.GetComponent("transform").(*Transform)
	rb := entity.GetComponent("rigidbody").(*Rigidbody)

	rb.Velocity.Multiply(rb.Friction)
	rb.Velocity.Add(rb.Gravity)
	t.Position.Add(pearl.Vector2Floor(rb.Velocity))
}

func (rbs *NonCollisionRigidbodySystem) Draw(entity *pearl.Entity, scene *pearl.Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions) {}

func (rbs *NonCollisionRigidbodySystem) GetRequirements() []string {
	return []string { "transform", "rigidbody" }
}