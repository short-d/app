package mdtest

import (
	"errors"

	"github.com/short-d/app/fw"
)

var _ fw.ProgramRuntime = (*ProgramRuntimeFake)(nil)

type ProgramRuntimeFake struct {
	callers []fw.Caller
}

func (p ProgramRuntimeFake) Caller(numLevelsUp int) (fw.Caller, error) {

	if numLevelsUp > len(p.callers) {
		return fw.Caller{}, errors.New("level of range")
	}
	return p.callers[numLevelsUp], nil
}

func (p ProgramRuntimeFake) LockOSThread() {
}

func NewProgramRuntimeFake(callers []fw.Caller) (ProgramRuntimeFake, error) {
	return ProgramRuntimeFake{callers: callers}, nil
}
