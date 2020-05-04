package timer

import (
	"time"
)

type Stub struct {
	CurrentTime time.Time
}

func (s Stub) Now() time.Time {
	return s.CurrentTime
}

func NewStub(currentTime time.Time) Stub {
	return Stub{
		CurrentTime: currentTime,
	}
}
