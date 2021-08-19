package builtInComponentsAndSystems


import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/eboatwright/pearl"
)


type onTriggerFunction func(other *pearl.Entity)


type Trigger struct {
	onTrigger onTriggerFunction
}
func (t *Trigger) ID() string { return "trigger" }


type TriggerSystem struct {}

func (ts *TriggerSystem) Update(entity *pearl.Entity, scene *pearl.Scene) {
	otherColliders := scene.FindEntitiesOfType("boxCollider")
	trigger := entity.GetComponent("trigger").(*Trigger)
	for _, other := range otherColliders {
		if EntitiesOverlap(entity, other) {
			trigger.onTrigger(other)
		}
	}
}

func (ts *TriggerSystem) Draw(entity *pearl.Entity, scene *pearl.Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions) {}

func (ts *TriggerSystem) GetRequirement() []string {
	return []string { "boxCollider", "trigger" }
}