package event

import "io"

type Dispatcher interface {
	Emitter
	Subscriber
	io.Closer
}
