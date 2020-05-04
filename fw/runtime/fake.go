package runtime

import "errors"

var _ Runtime = (*Fake)(nil)

type Fake struct {
	callers []Caller
}

func (f Fake) Caller(numLevelsUp int) (Caller, error) {
	if numLevelsUp > len(f.callers) {
		return Caller{}, errors.New("level of range")
	}
	return f.callers[numLevelsUp], nil
}

func (f Fake) LockOSThread() {
}

func NewFake(callers []Caller) (Fake, error) {
	return Fake{callers: callers}, nil
}
