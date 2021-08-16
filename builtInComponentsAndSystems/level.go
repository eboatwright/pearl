package builtInComponentsAndSystems


import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/eboatwright/pearl"
)


type Level struct {
	Data     [][]int
	Tileset  *ebiten.Image
	TileSize int
}
func (l *Level) ID() string { return "level" }


type LevelRenderer struct {}

func (lr *LevelRenderer) Update(entity *pearl.Entity, scene *pearl.Scene) {}

func (lr *LevelRenderer) Draw(entity *pearl.Entity, scene *pearl.Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	l := entity.GetComponent("level").(*Level)

	for y := 0; y < len(l.Data); y++ {
		for x := 0; x < len(l.Data[0]); x++ {
		}
	}
}

func (lr *LevelRenderer) GetRequirements() []string {
	return []string { "level" }
}