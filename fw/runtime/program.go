package runtime

import (
	"errors"
	"runtime"
)

var _ Runtime = (*Program)(nil)

type Program struct {
}

func (p Program) LockOSThread() {
	runtime.LockOSThread()
}

func (p Program) Caller(numLevelsUp int) (Caller, error) {
	_, file, line, ok := runtime.Caller(numLevelsUp + 1)
	if !ok {
		return Caller{}, errors.New("fail to obtain caller info")
	}
	return Caller{
		FullFilename: file,
		LineNumber:   line,
	}, nil
}

func NewProgram() Program {
	return Program{}
}
