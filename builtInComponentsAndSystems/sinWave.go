package builtInComponentsAndSystems


import (
	"github.com/eboatwright/pearl"
	"math"
)


type SinWave struct {
	Offset   float64
	Speed    float64
	Distance float64
}
func (sw *SinWave) ID() string { return "sinWave" }

func (sw *SinWave) GetSin() float64 {
	return math.Sin(sw.Speed * pearl.TimeSinceStart()) * sw.Distance
}