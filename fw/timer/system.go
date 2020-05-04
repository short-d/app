package timer

import "time"

var _ Timer = (*System)(nil)

type System struct{}

func (t System) Now() time.Time {
	return time.Now()
}

func NewSystem() System {
	return System{}
}
