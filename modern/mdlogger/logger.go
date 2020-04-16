package mdlogger

import (
	"fmt"

	"github.com/short-d/app/fw"
)

var _ fw.Logger = (*Logger)(nil)

type priority int

var priorities = map[fw.LogLevel]priority{
	fw.LogOff:   0,
	fw.LogFatal: 1,
	fw.LogError: 2,
	fw.LogWarn:  3,
	fw.LogInfo:  4,
	fw.LogDebug: 5,
	fw.LogTrace: 6,
}

type Logger struct {
	prefix         string
	level          fw.LogLevel
	timer          fw.Timer
	programRuntime fw.ProgramRuntime
	entryRepo      EntryRepository
}

func (l Logger) Fatal(message string) {
	if l.levelAbove(fw.LogFatal) {
		return
	}
	l.log(fw.LogFatal, message)
}

func (l Logger) Error(err error) {
	if l.levelAbove(fw.LogError) {
		return
	}
	l.log(fw.LogError, fmt.Sprintf("%v", err))
}

func (l Logger) Warn(message string) {
	if l.levelAbove(fw.LogWarn) {
		return
	}
	l.log(fw.LogWarn, message)
}

func (l Logger) Info(message string) {
	if l.levelAbove(fw.LogInfo) {
		return
	}
	l.log(fw.LogInfo, message)
}

func (l Logger) Debug(message string) {
	if l.levelAbove(fw.LogDebug) {
		return
	}
	l.log(fw.LogDebug, message)
}

func (l Logger) Trace(message string) {
	if l.levelAbove(fw.LogTrace) {
		return
	}
	l.log(fw.LogTrace, message)
}

func (l Logger) log(level fw.LogLevel, message string) {
	now := l.timer.Now().UTC()
	caller, err := l.programRuntime.Caller(2)
	if err != nil {
		l.entryRepo.createLogEntry(level, l.prefix, 0, "", message, now)
		return
	}
	l.entryRepo.createLogEntry(level, l.prefix, caller.LineNumber, caller.FullFilename, message, now)
}

func (l Logger) levelAbove(logLevel fw.LogLevel) bool {
	return priorities[l.level] < priorities[logLevel]
}

func NewLogger(
	prefix string,
	level fw.LogLevel,
	timer fw.Timer,
	programRuntime fw.ProgramRuntime,
	entryRepo EntryRepository,
) Logger {
	return Logger{
		prefix:         prefix,
		level:          level,
		timer:          timer,
		programRuntime: programRuntime,
		entryRepo:      entryRepo,
	}
}
