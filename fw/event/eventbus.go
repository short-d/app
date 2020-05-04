package event

import (
	"sync"

	"github.com/asaskevich/EventBus"
)

var _ Dispatcher = (*Eventbus)(nil)

// EventBus publishes an event to all its subscribers
// TODO(issue#84): replace Eventbus with github.com/short-d/eventbus
type Eventbus struct {
	eventBus EventBus.Bus
	lock     sync.RWMutex
	isClosed bool
}

func (d *Eventbus) Dispatch(event Event) error {
	d.lock.RLock()
	defer d.lock.RUnlock()

	if d.isClosed {
		return ErrDispatcherIsClosed
	}

	d.eventBus.Publish(event.GetName(), event)

	return nil
}

func (d *Eventbus) Subscribe(listener Listener) error {
	d.lock.RLock()
	defer d.lock.RUnlock()

	if d.isClosed {
		return ErrDispatcherIsClosed
	}

	return d.eventBus.SubscribeAsync(listener.GetSubscribedEvent(), listener.Handle, false)
}

func (d *Eventbus) Unsubscribe(listener Listener) error {
	d.lock.RLock()
	defer d.lock.RUnlock()

	if d.isClosed {
		return ErrDispatcherIsClosed
	}

	return d.eventBus.Unsubscribe(listener.GetSubscribedEvent(), listener.Handle)
}

func (d *Eventbus) BindListeners(listeners []Listener) error {
	for _, listener := range listeners {
		if err := d.Subscribe(listener); err != nil {
			return err
		}
	}

	return nil
}

func (d *Eventbus) Close() error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.isClosed {
		return ErrDispatcherIsClosed
	}

	d.isClosed = true
	d.eventBus.WaitAsync()

	return nil
}

// NewEventBus creates EventDispatcher
func NewEventBus() *Eventbus {
	return &Eventbus{
		eventBus: EventBus.New(),
		lock:     sync.RWMutex{},
		isClosed: false,
	}
}
