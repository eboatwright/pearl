package pearl


import (
    "github.com/hajimehoshi/ebiten/v2"
)


// The struct that all Scenes are
type Scene struct {
    // The Scene's name / ID
	ID       string
    // A list of the Scene's Entities
    Entities []*Entity
    // A list of the Scene's Systems
    Systems  []ISystem
}

// Adds Entity specified to Scene's entity list
func (s *Scene) AddEntity(entity *Entity) {
	s.Entities = append(s.Entities, entity)
}

// Adds Entities specified to Scene's entity list
func (s *Scene) AddEntities(entities []*Entity) {
    for _, entity := range entities {
        s.AddEntity(entity)
    }
}

// Removes Entity at index (might cause problems if you use this. It's best to just call Destroy() and let Pearl handle it, but you do you XD)
func (s *Scene) RemoveEntity(index int) {
    s.Entities = append(s.Entities[:index], s.Entities[index + 1:]...)
}

// Returns Entity from Scene's entity list with ID specified, and returns nil if Scene doesn't have an Entity with that ID
func (s *Scene) GetEntityWithID(id string) *Entity {
    for _, entity := range s.Entities {
        if entity.ID == id {
            return entity
        }
    }
    return nil
}

// Returns first Entity found with tag, but returns nil if not found
func (s *Scene) GetEntityWithTag(tag string) *Entity {
    for _, entity := range s.Entities {
        if entity.HasTag(tag) {
            return entity
        }
    }
    return nil
}

// Returns a slice of Entities found with the specified tag
func (s *Scene) GetEntitiesWithTag(tag string) []*Entity {
    found := []*Entity {  }
    for _, entity := range s.Entities {
        if entity.HasTag(tag) {
            found = append(found, entity)
        }
    }
    return found
}

// Returns the first Entity found with specified component, but if none found returns nil
func (s *Scene) FindEntityOfType(id string) *Entity {
    for _, entity := range s.Entities {
        if entity.HasComponent(id) {
            return entity
        }
    }
    return nil
}

// Returns a slice of Entities found with Component
func (s *Scene) FindEntitiesOfType(id string) []*Entity {
    found := []*Entity {  }
    for _, entity := range s.Entities {
        if entity.HasComponent(id) {
            found = append(found, entity)
        }
    }
    return found
}

// Returns the first component found on the first entity found, returns nil if not found
func (s *Scene) FindComponent(id string) IComponent {
    for _, entity := range s.Entities {
        component := entity.GetComponent(id)
        if component != nil {
            return component
        }
    }
    return nil
}

// Returns a slice of Components found, returns empty slice if none found
func (s *Scene) FindComponents(id string) []IComponent {
    found := []IComponent {  }
    for _, entity := range s.Entities {
        component := entity.GetComponent(id)
        if component != nil {
            found = append(found, component)
        }
    }
    return found
}

// Adds System specified to Scene's System list
func (s *Scene) AddSystem(system ISystem) {
    s.Systems = append(s.Systems, system)
}

// Adds Systems specified to Scene's System list
func (s *Scene) AddSystems(systems []ISystem) {
    for _, system := range systems {
        s.AddSystem(system)
    }
}

// This is automatically called, and it updates all Entities and Systems
func (s *Scene) Update() {
    for _, system := range s.Systems {
        for i := len(s.Entities) - 1; i >= 0; i-- {
            if s.Entities[i].Destroyed {
                s.RemoveEntity(i)
            } else if EntityMatchesSystem(s.Entities[i], system) {
                system.Update(s.Entities[i], s)
            }
        }
    }
}

// This is automatically called, and it draws all Entities and Systems
func (s *Scene) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions {}
    
	for _, system := range s.Systems {
        for _, entity := range s.Entities {
            if EntityMatchesSystem(entity, system) {
                system.Draw(entity, s, screen, options)
            }
        }
    }
}