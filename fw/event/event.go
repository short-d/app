package event

import (
	"errors"
)

// ErrDispatcherIsClosed represents that there is no way to perform manipulations with event dispatcher
var ErrDispatcherIsClosed = errors.New("failed to perform the operation, the dispatcher is closed")

// Event represents a message that expects to be delivered to the corresponding listeners
type Event interface {
	// GetName returns the name of an event
	GetName() string
}
