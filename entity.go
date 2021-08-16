package pearl


// The struct that all entities are
type Entity struct {
	// The name / ID of the entity
	ID         string
	// Whether or not the Entity is destroyed
	Destroyed  bool
	components []IComponent
	tags       []string
	parent     *Entity
	children   []*Entity
}

// Flags entity as destroyed and will destroy it in the next update loop
func (e *Entity) Destroy() {
	e.Destroyed = true
}

// Adds component specified to entity's component list
func (e *Entity) AddComponent(component IComponent) {
	e.components = append(e.components, component)
}

// Adds each of the components specified to entity's component list
func (e *Entity) AddComponents(components []IComponent) {
	for _, component := range components {
		e.AddComponent(component)
	}
}

// Returns the component with the id specified, but returns nil if entity doesn't a component with that id
func (e *Entity) GetComponent(id string) IComponent {
	if !e.HasComponent(id) { return nil }
	for _, component := range e.components {
		if component.ID() == id {
			return component
		}
	}
	return nil
}

// Returns a slice of components found on entity, returns an empty slice if none found
func (e *Entity) GetComponents(id string) []IComponent {
	found := []IComponent {}
	for _, component := range e.components {
		if component.ID() == id {
			found = append(found, component)
		}
	}
	return found
}

// Returns the component in the parent with id specified. Returns nil if parent doesn't have component
func (e *Entity) GetComponentInParent(id string) IComponent {
	return e.parent.GetComponent(id)
}

// Returns a slice of components with id specified found in the parent. Returns empty list if none found
func (e *Entity) GetComponentsInParent(id string) []IComponent {
	return e.parent.GetComponents(id)
}

// Checks each child for component with id and returns the first one found. Returns nil if not found
func (e *Entity) GetComponentInChildren(id string) IComponent {
	for _, child := range e.children {
		if child.HasComponent(id) {
			return child.GetComponent(id)
		}
	}
	return nil
}

// Returns slice of all components with id found in each children, returns empty slice if none found
func (e *Entity) GetComponentsInChildren(id string) []IComponent {
	found := []IComponent {}
	for _, child := range e.children {
		for _, component := range child.components {
			if component.ID() == id {
				found = append(found, component)
			}
		}
	}
	return found
}

// Returns bool on wether or not the entity has the component
func (e *Entity) HasComponent(id string) bool {
	for _, component := range e.components {
		if component.ID() == id {
			return true
		}
	}
	return false
}

// Adds tag to the entity's tag list
func (e *Entity) AddTag(tag string) {
	e.tags = append(e.tags, tag)
}

// Adds each of the tags specified to entity's tag list
func (e *Entity) AddTags(tags []string) {
	for _, tag := range tags {
		e.AddTag(tag)
	}
}

// Returns whether or not the Entity has the tag specified
func (e *Entity) HasTag(tag string) bool {
	for _, t := range e.tags {
		if t == tag {
			return true
		}
	}
	return false
}

// Sets Entity's parent variable
func (e *Entity) SetParent(entity *Entity) {
	e.parent = entity
}

// Returns Entity's parent
func (e *Entity) GetParent() *Entity {
	return e.parent
}

// Adds Entity to child list
func (e *Entity) AddChild(entity *Entity) {
	e.children = append(e.children, entity)
}

// Adds Entities to child list
func (e *Entity) AddChildren(entities []*Entity) {
	for _, entity := range entities {
		e.AddChild(entity)
	}
}

// Removes Entity at index from child list, but does nothing if entity specified isn't a child of entity
func (e *Entity) RemoveChildAt(index int) {
	e.children = append(e.children[:index], e.children[index + 1:]...)
}

// Removes Entity from child list, but does nothing if entity specified isn't a child of entity
func (e *Entity) RemoveChild(entity *Entity) {
	for i := 0; i < len(e.children); i++ {
		if e.children[i] == entity {
			e.RemoveChildAt(i)
			break
		}
	}
}

// Returns Entity's children
func (e *Entity) GetChildren() []*Entity {
	return e.children
}