package timer

import (
	"time"
)

type Timer interface {
	Now() time.Time
	Ticker(interval time.Duration, operation func()) chan bool
}
