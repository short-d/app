package event

// Listener handles received events
type Listener interface {
	// Handle processes the received event
	Handle(event Event)
	// GetSubscribedEvent returns the event name this listener wants to listen to
	GetSubscribedEvent() string
}
