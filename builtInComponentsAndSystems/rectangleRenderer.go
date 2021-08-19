package builtInComponentsAndSystems


import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/eboatwright/pearl"

	"image/color"
)


type RectangleRenderer struct {
	Color color.RGBA
}
func (rr *RectangleRenderer) ID() string { return "rectangleRenderer" }


type RectangleRendererSystem struct {}

func (rrs *RectangleRendererSystem) Update(entity *pearl.Entity, scene *pearl.Scene) {}

func (rrs *RectangleRendererSystem) Draw(entity *pearl.Entity, scene *pearl.Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	t := entity.GetComponent("transform").(*Transform)
	rr := entity.GetComponent("rectangleRenderer").(*RectangleRenderer)
	
	ebitenutil.DrawRect(
		screen,
		t.Position.X, t.Position.Y,
		t.Scale.X, t.Scale.Y,
		rr.Color,
	)
}

func (rrs *RectangleRendererSystem) GetRequirements() []string {
	return []string { "transform", "rectangleRenderer" }
}