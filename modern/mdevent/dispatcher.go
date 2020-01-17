package mdevent

import (
	"sync"

	"github.com/short-d/app/fw"

	"github.com/asaskevich/EventBus"
)

var _ fw.Dispatcher = (*EventDispatcher)(nil)

// EventDispatcher publishes an event to all its subscribers
type EventDispatcher struct {
	eventBus EventBus.Bus
	lock     sync.RWMutex
	isClosed bool
}

func (d *EventDispatcher) Dispatch(event fw.Event) error {
	d.lock.RLock()
	defer d.lock.RUnlock()

	if d.isClosed {
		return fw.ErrDispatcherIsClosed
	}

	d.eventBus.Publish(event.GetName(), event)

	return nil
}

func (d *EventDispatcher) Subscribe(listener fw.Listener) error {
	d.lock.RLock()
	defer d.lock.RUnlock()

	if d.isClosed {
		return fw.ErrDispatcherIsClosed
	}

	return d.eventBus.SubscribeAsync(listener.GetSubscribedEvent(), listener.Handle, false)
}

func (d *EventDispatcher) Unsubscribe(listener fw.Listener) error {
	d.lock.RLock()
	defer d.lock.RUnlock()

	if d.isClosed {
		return fw.ErrDispatcherIsClosed
	}

	return d.eventBus.Unsubscribe(listener.GetSubscribedEvent(), listener.Handle)
}

func (d *EventDispatcher) BindListeners(listeners []fw.Listener) error {
	for _, listener := range listeners {
		if err := d.Subscribe(listener); err != nil {
			return err
		}
	}

	return nil
}

func (d *EventDispatcher) Close() error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.isClosed {
		return fw.ErrDispatcherIsClosed
	}

	d.isClosed = true
	d.eventBus.WaitAsync()

	return nil
}

// NewEventDispatcher creates a new instance of EventDispatcher type
func NewEventDispatcher(eventBus EventBus.Bus) *EventDispatcher {
	return &EventDispatcher{
		eventBus: eventBus,
		lock:     sync.RWMutex{},
		isClosed: false,
	}
}
