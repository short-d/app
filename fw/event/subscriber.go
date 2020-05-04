package event

// Subscriber provides the ability to subscribe a listener to an event
type Subscriber interface {
	// Subscribe binds the given listener to events with the provided event name
	Subscribe(listener Listener) error
	// Unsubscribe unbinds the given listener from the provided event name
	Unsubscribe(listener Listener) error
	// BindListeners registers the list of listeners
	BindListeners(listeners []Listener) error
}
