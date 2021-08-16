package builtInComponentsAndSystems


import (
	"github.com/eboatwright/pearl"
)


type Transform struct {
	Position pearl.Vector2
	Scale    pearl.Vector2
	Rotation float64
}
func (t *Transform) ID() string { return "transform" }