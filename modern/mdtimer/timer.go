package mdtimer

import (
	"time"

	"github.com/byliuyang/app/fw"
)

type Timer struct{}

func (t Timer) Now() time.Time {
	return time.Now()
}

func NewTimer() fw.Timer {
	return Timer{}
}
