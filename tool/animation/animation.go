package animation

import (
	"time"

	"github.com/byliuyang/app/tool/ticker"
)

type Animation struct {
	frames       []string
	interval     time.Duration
	currFrameIdx int
	frameCount   int
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
	a.done = ticker.NewTicker(a.interval, func() {
		a.nextFrame()
	})
}

func (a *Animation) Stop() {
	a.done <- true
}

func NewAnimation(frames []string, interval time.Duration) Animation {
	return Animation{
		frames:       frames,
		frameCount:   len(frames),
		interval:     interval,
		currFrameIdx: 0,
	}
}
