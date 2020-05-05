package logger

import (
	"time"

	"github.com/short-d/app/fw/runtime"
	"github.com/short-d/app/fw/timer"
)

func NewFake(level LogLevel, entryRepo EntryRepository) (Logger, error) {
	tm := timer.NewStub(time.Now())
	rt, err := runtime.NewFake([]runtime.Caller{})
	if err != nil {
		return Logger{}, err
	}
	return NewLogger("Fake", level, tm, rt, entryRepo), nil
}
