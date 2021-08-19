package builtInComponentsAndSystems


import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/eboatwright/pearl"
)


type Animation struct {
	Frames []int
	Speed  int
}

type Animator struct {
	Animations            []Animation
	CurrentAnimation      int
	CurrentAnimationIndex int
	AnimationTimer        int
}
func (a *Animator) ID() string { return "animator" }

func (a *Animator) ChangeAnimation(index int) {
	if a.CurrentAnimation != index {
		a.CurrentAnimation = index
		a.CurrentAnimationIndex = 0
		a.AnimationTimer = 0
	}
}


type AnimationSystem struct {}

func (as *AnimationSystem) Update(entity *pearl.Entity, scene *pearl.Scene) {
	a := entity.GetComponent("animator").(*Animator)
	i := entity.GetComponent("image").(*Image)

	a.AnimationTimer--
	if a.AnimationTimer <= 0 {
		a.AnimationTimer = a.Animations[a.CurrentAnimation].Speed
		a.CurrentAnimationIndex++

		if a.CurrentAnimationIndex >= len(a.Animations[a.CurrentAnimation].Frames) {
			a.CurrentAnimationIndex = 0
		}

		i.SourcePos.X = float64(a.Animations[a.CurrentAnimation].Frames[a.CurrentAnimationIndex]) * i.Size.X
	}
}

func (as *AnimationSystem) Draw(entity *pearl.Entity, scene *pearl.Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions) {}

func (as *AnimationSystem) GetRequirements() []string {
	return []string { "animator", "image" }
}