package service

import (
	"github.com/short-d/app/fw/io"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/runtime"
	"github.com/short-d/app/fw/timer"
)

func newDefaultLogger(name string) logger.Logger {
	tm := timer.NewSystem()
	rt := runtime.NewProgram()
	stdOut := io.NewStdOut()
	entryRepo := logger.NewLocal(stdOut, false)
	return logger.NewLogger(name, logger.LogInfo, tm, rt, entryRepo)
}
