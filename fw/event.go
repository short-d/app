package fw

import (
	"errors"
	"io"
)

// ErrDispatcherIsClosed represents that there is no way to perform manipulations with event dispatcher
var ErrDispatcherIsClosed = errors.New("failed to perform the operation, the dispatcher is closed")

type Dispatcher interface {
	Emitter
	Subscriber
	io.Closer
}

// Emitter propagates the given event to the list of corresponding listeners
type Emitter interface {
	// Dispatch dispatches an event to all registered listeners
	Dispatch(event Event) error
}

// Listener handles received events
type Listener interface {
	// Handle processes the received event
	Handle(event Event)
	// GetSubscribedEvent returns the event name this listener wants to listen to
	GetSubscribedEvent() string
}

// Event represents a message that expects to be delivered to the corresponding listeners
type Event interface {
	// GetName returns the name of an event
	GetName() string
}

// Subscriber provides the ability to subscribe a listener to an event
type Subscriber interface {
	// Subscribe binds the given listener to events with the provided event name
	Subscribe(listener Listener) error
	// Unsubscribe unbinds the given listener from the provided event name
	Unsubscribe(listener Listener) error
	// BindListeners registers the list of listeners
	BindListeners(listeners []Listener) error
}
