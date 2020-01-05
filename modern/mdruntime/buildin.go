package mdruntime

import (
	"errors"
	"runtime"

	"github.com/short-d/app/fw"
)

var _ fw.ProgramRuntime = (*BuildIn)(nil)

type BuildIn struct {
}

func (b BuildIn) Caller(numLevelsUp int) (fw.Caller, error) {
	_, file, line, ok := runtime.Caller(numLevelsUp)
	if !ok {
		return fw.Caller{}, errors.New("fail to obtain caller info")
	}
	return fw.Caller{
		FullFilename: file,
		LineNumber:   line,
	}, nil
}

func NewBuildIn() BuildIn {
	return BuildIn{}
}