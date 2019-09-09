package mdtest

import (
	"time"
)

type Timer struct {
	CurrentTime time.Time
}

func (t Timer) Now() time.Time {
	return t.CurrentTime
}

func NewTimerFake(currentTime time.Time) Timer {
	return Timer{
		CurrentTime: currentTime,
	}
}
