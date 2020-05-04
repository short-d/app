package timer

import (
	"time"
)

type Timer interface {
	Now() time.Time
}
