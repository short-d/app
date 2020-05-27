package ui

import (
	"time"

	"github.com/short-d/app/fw/animation"
	"github.com/short-d/app/fw/timer"
)

func NewSpinner(timer timer.Timer) animation.Animation {
	return animation.NewAnimation(
		[]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		60*time.Millisecond,
		timer,
	)
}
