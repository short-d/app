package timer

import "time"

var _ Timer = (*System)(nil)

type System struct{}

func (t System) Now() time.Time {
	return time.Now()
}

func (t System) Ticker(interval time.Duration, operation func()) chan bool {
	done := make(chan bool)
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				operation()
			}
		}
	}()
	return done
}

func NewSystem() System {
	return System{}
}
