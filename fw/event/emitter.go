package event

// Emitter propagates the given event to the list of corresponding listeners
type Emitter interface {
	// Dispatch dispatches an event to all registered listeners
	Dispatch(event Event) error
}
