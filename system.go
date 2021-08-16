package pearl


import (
    "github.com/hajimehoshi/ebiten/v2"
)


// The interface all systems will implement
type ISystem interface {
    // Your update loop (called 60 times per second) before draw
    Update(entity *Entity, scene *Scene)
    // Your draw loop (called 60 times per second) after update
    Draw(entity *Entity, scene *Scene, screen *ebiten.Image, options *ebiten.DrawImageOptions)
    // Should return a slice of strings that have the IDs of all components required
    GetRequirements() []string
}


// Returns a bool on whether or not an entity meets the requirements of the system
func EntityMatchesSystem(entity *Entity, system ISystem) bool {
    for _, requirement := range system.GetRequirements() {
        if !entity.HasComponent(requirement) {
            return false
        }
    }
    return true
}