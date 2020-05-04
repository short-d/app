package animation

import (
	"time"

	"github.com/short-d/app/fw/timer"
)

type Animation struct {
	frames       []string
	interval     time.Duration
	currFrameIdx int
	frameCount   int
	timer        timer.Timer
	done         chan bool
}

func (a Animation) Draw() string {
	frame := a.frames[a.currFrameIdx]
	return frame
}

func (a *Animation) nextFrame() {
	a.currFrameIdx = (a.currFrameIdx + 1) % a.frameCount
}

func (a *Animation) Start() {
	a.done = a.timer.Ticker(a.interval, func() {
		a.nextFrame()
	})
}

func (a *Animation) Stop() {
	a.done <- true
}

func NewAnimation(frames []string, interval time.Duration, timer timer.Timer) Animation {
	return Animation{
		frames:       frames,
		frameCount:   len(frames),
		interval:     interval,
		currFrameIdx: 0,
		timer:        timer,
	}
}
