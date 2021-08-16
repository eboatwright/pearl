package builtInComponentsAndSystems


import (
	"github.com/hajimehoshi/ebiten/v2"
    "github.com/eboatwright/pearl"

	"math"
	"image"
)


type Image struct {
	Image     *ebiten.Image
    Size      pearl.Vector2
    SourcePos pearl.Vector2
}
func (t *Image) ID() string { return "image" }


type ImageRenderer struct {}

func (ir *ImageRenderer) Update(entity *pearl.Entity, scene *pearl.Scene) {}

func (ir *ImageRenderer) Draw(entity *pearl.Entity, scene *pearl.Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions) {
    t := entity.GetComponent("transform").(*Transform)
    i := entity.GetComponent("image").(*Image)

	options.GeoM.Reset()
	options.GeoM.Scale(t.Scale.X, t.Scale.Y)
	options.GeoM.Translate(
		-(i.Size.X * t.Scale.X) / 2,
		-(i.Size.Y * t.Scale.Y) / 2,
	)
	options.GeoM.Rotate(t.Rotation * math.Pi / 180)
	options.GeoM.Translate(
		t.Position.X,
		t.Position.Y,
	)

	screen.DrawImage(i.Image.SubImage(image.Rect(int(i.SourcePos.X), int(i.SourcePos.Y), int(i.SourcePos.X + i.Size.X), int(i.SourcePos.Y + i.Size.Y))).(*ebiten.Image), options)
}

func (ir *ImageRenderer) GetRequirements() []string {
    return []string { "transform", "image" }
}