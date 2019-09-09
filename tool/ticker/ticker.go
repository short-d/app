package ticker

import "time"

func NewTicker(interval time.Duration, operation func()) chan bool {
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
