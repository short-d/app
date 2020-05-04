package timer

import (
	"time"
)

var _ Timer = (*Stub)(nil)

type Stub struct {
	CurrentTime time.Time
}

func (s Stub) Ticker(interval time.Duration, operation func()) chan bool {
	return make(chan bool)
}

func (s Stub) Now() time.Time {
	return s.CurrentTime
}

func NewStub(currentTime time.Time) Stub {
	return Stub{
		CurrentTime: currentTime,
	}
}
