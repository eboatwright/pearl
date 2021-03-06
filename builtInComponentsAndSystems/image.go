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
	Centered  bool
}
func (t *Image) ID() string { return "image" }


type ImageRenderer struct {}

func (ir *ImageRenderer) Update(entity *pearl.Entity, scene *pearl.Scene) {}

func (ir *ImageRenderer) Draw(entity *pearl.Entity, scene *pearl.Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions) {
    t := entity.GetComponent("transform").(*Transform)
    i := entity.GetComponent("image").(*Image)

	options.GeoM.Reset()
	options.GeoM.Scale(t.Scale.X, t.Scale.Y)
	
	if i.Centered {
		options.GeoM.Translate(
			-i.Size.X / 2,
			-i.Size.Y / 2,
		)
	} else {
		if t.Scale.X < 0 {
			options.GeoM.Translate(-(i.Size.X * t.Scale.X), 0)
		}
		if t.Scale.Y < 0 {
			options.GeoM.Translate(0, -(i.Size.Y * t.Scale.Y))
		}
	}

	options.GeoM.Rotate(t.Rotation * math.Pi / 180)

	parent := entity.GetParent()
	if parent != nil {
		parentT := parent.GetComponent("transform").(*Transform)
		if parentT != nil {
			options.GeoM.Rotate(parentT.Rotation * math.Pi / 180)
			options.GeoM.Translate(
				parentT.Position.X,
				parentT.Position.Y,
			)
		}
	}
	options.GeoM.Translate(
		t.Position.X,
		t.Position.Y,
	)
	if entity.HasComponent("sinWave") {
		sw := entity.GetComponent("sinWave").(*SinWave)
		options.GeoM.Translate(
			0,
			sw.GetSin(),
		)
	}

	screen.DrawImage(i.Image.SubImage(image.Rect(int(i.SourcePos.X), int(i.SourcePos.Y), int(i.SourcePos.X + i.Size.X), int(i.SourcePos.Y + i.Size.Y))).(*ebiten.Image), options)
}

func (ir *ImageRenderer) GetRequirements() []string {
    return []string { "transform", "image" }
}