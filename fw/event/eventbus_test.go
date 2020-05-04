package event

import (
	"sync/atomic"
	"testing"

	"github.com/short-d/app/fw/assert"
)

func TestEventbus_Dispatch(t *testing.T) {
	testCases := []struct {
		name      string
		events    []Event
		listeners []Listener
	}{
		{
			name: "one event one listener",
			events: []Event{
				fakeEvent("event1"),
			},
			listeners: []Listener{
				&fakeListener{
					name:          "event1",
					expectedCalls: 1,
				},
			},
		},
		{
			name: "one event two listener",
			events: []Event{
				fakeEvent("event1"),
			},
			listeners: []Listener{
				&fakeListener{
					name:          "event1",
					expectedCalls: 1,
				},
				&fakeListener{
					name:          "event1",
					expectedCalls: 1,
				},
			},
		},
		{
			name: "one of the listeners listens to another event",
			events: []Event{
				fakeEvent("event1"),
			},
			listeners: []Listener{
				&fakeListener{
					name:          "event1",
					expectedCalls: 1,
				},
				&fakeListener{
					name:          "event2",
					expectedCalls: 0,
				},
			},
		},
		{
			name: "two events two listeners",
			events: []Event{
				fakeEvent("event1"),
				fakeEvent("event2"),
			},
			listeners: []Listener{
				&fakeListener{
					name:          "event2",
					expectedCalls: 1,
				},
				&fakeListener{
					name:          "event1",
					expectedCalls: 1,
				},
			},
		},
		{
			name: "multiple events calls",
			events: []Event{
				fakeEvent("event1"),
				fakeEvent("event1"),
				fakeEvent("event2"),
				fakeEvent("event1"),
				fakeEvent("event1"),
				fakeEvent("event2"),
			},
			listeners: []Listener{
				&fakeListener{
					name:          "event2",
					expectedCalls: 2,
				},
				&fakeListener{
					name:          "event1",
					expectedCalls: 4,
				},
				&fakeListener{
					name:          "event1",
					expectedCalls: 4,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			Eventbus := NewEventBus()

			err := Eventbus.BindListeners(testCase.listeners)
			assert.Equal(t, nil, err)

			for _, e := range testCase.events {
				err = Eventbus.Dispatch(e)
				assert.Equal(t, nil, err)
			}

			err = Eventbus.Close()
			assert.Equal(t, nil, err)

			for _, obj := range testCase.listeners {
				listener, ok := obj.(*fakeListener)

				assert.Equal(t, true, ok)
				assert.Equal(t, listener.expectedCalls, listener.actualCalls)
			}
		})
	}
}

func TestEventbus_Close(t *testing.T) {
	ev := fakeEvent("event1")
	Eventbus := NewEventBus()
	listener := &fakeListener{name: ev.GetName()}

	_ = Eventbus.BindListeners([]Listener{listener})

	err := Eventbus.Dispatch(ev)
	assert.Equal(t, nil, err)

	err = Eventbus.Close()
	assert.Equal(t, nil, err)

	err = Eventbus.Dispatch(ev)
	assert.Equal(t, ErrDispatcherIsClosed, err)

	// there was no any listener call, because we have unsubscribed the listener
	assert.Equal(t, int32(1), listener.actualCalls)
	assert.Equal(t, ErrDispatcherIsClosed, Eventbus.Close())
}

type fakeListener struct {
	name          string
	expectedCalls int32
	actualCalls   int32
}

func (f *fakeListener) Handle(e Event) {
	// atomically increase actualCalls number, because of concurrent access to this variable
	atomic.AddInt32(&f.actualCalls, 1)
}

func (f *fakeListener) GetSubscribedEvent() string {
	return f.name
}

type fakeEvent string

func (f fakeEvent) GetName() string {
	return string(f)
}
