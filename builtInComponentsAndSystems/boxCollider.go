package builtInComponentsAndSystems


import (
	"github.com/eboatwright/pearl"
)


type BoxCollider struct {
	Size   pearl.Vector2
	Offset pearl.Vector2
}
func (bc *BoxCollider) ID() string { return "boxCollider" }


func EntitiesOverlap(a *pearl.Entity, b *pearl.Entity) bool {
	aPosition := a.GetComponent("transform").(*Transform).Position
	bPosition := b.GetComponent("transform").(*Transform).Position
	aBc := a.GetComponent("boxCollider").(*BoxCollider)
	bBc := b.GetComponent("boxCollider").(*BoxCollider)

	aPos := pearl.Vector2Add(aPosition, aBc.Offset)
	bPos := pearl.Vector2Add(bPosition, bBc.Offset)

	return aPos.X < bPos.X + bBc.Size.X &&
		   aPos.X + aBc.Size.X > bPos.X &&
		   aPos.Y < bPos.Y + bBc.Size.Y &&
		   aPos.Y + aBc.Size.Y > bPos.Y
}