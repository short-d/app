package logger

import (
	"fmt"

	"github.com/short-d/app/fw"
)

type LogLevel int

const (
	LogFatal LogLevel = iota
	LogError
	LogWarn
	LogInfo
	LogDebug
	LogTrace
	LogOff
)

type priority int

var priorities = map[LogLevel]priority{
	LogOff:   0,
	LogFatal: 1,
	LogError: 2,
	LogWarn:  3,
	LogInfo:  4,
	LogDebug: 5,
	LogTrace: 6,
}

type Logger struct {
	prefix         string
	level          LogLevel
	timer          fw.Timer
	programRuntime fw.ProgramRuntime
	entryRepo      EntryRepository
}

func (l Logger) Fatal(message string) {
	if l.levelAbove(LogFatal) {
		return
	}
	l.log(LogFatal, message)
}

func (l Logger) Error(err error) {
	if l.levelAbove(LogError) {
		return
	}
	l.log(LogError, fmt.Sprintf("%v", err))
}

func (l Logger) Warn(message string) {
	if l.levelAbove(LogWarn) {
		return
	}
	l.log(LogWarn, message)
}

func (l Logger) Info(message string) {
	if l.levelAbove(LogInfo) {
		return
	}
	l.log(LogInfo, message)
}

func (l Logger) Debug(message string) {
	if l.levelAbove(LogDebug) {
		return
	}
	l.log(LogDebug, message)
}

func (l Logger) Trace(message string) {
	if l.levelAbove(LogTrace) {
		return
	}
	l.log(LogTrace, message)
}

func (l Logger) log(level LogLevel, message string) {
	now := l.timer.Now().UTC()
	caller, err := l.programRuntime.Caller(2)
	if err != nil {
		l.entryRepo.createLogEntry(level, l.prefix, 0, "", message, now)
		return
	}
	l.entryRepo.createLogEntry(level, l.prefix, caller.LineNumber, caller.FullFilename, message, now)
}

func (l Logger) levelAbove(logLevel LogLevel) bool {
	return priorities[l.level] < priorities[logLevel]
}

func NewLogger(
	prefix string,
	level LogLevel,
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
