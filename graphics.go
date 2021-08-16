package pearl


import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"image/color"
	_ "image/png"
)


// Returns a color.RGBA with the R, G, B, A provided
func RGBA(r, g, b, a uint8) color.RGBA {
	return color.RGBA { r, g, b, a }
}


// A helper function for getting an Ebiten image
func LoadImage(path string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile(path)

	if err != nil {
		panic(err)
	}

	return image
}

// Helper function for setting a ebiten.DrawImageOptions.GeoM position (resets GeoM before doing so because Ebiten ¯\_(ツ)_/¯)
func SetDrawPosition(options *ebiten.DrawImageOptions, position Vector2) {
	options.GeoM.Reset()
	options.GeoM.Translate(position.X, position.Y)
}