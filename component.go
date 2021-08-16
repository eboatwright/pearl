package pearl


// The interface all components will implement
type IComponent interface {
	// Returns name / ID of component
	ID() string
}