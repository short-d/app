package animation

import (
	"time"
)

func NewSpinner() Animation {
	return NewAnimation(
		[]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		60*time.Millisecond,
	)
}
