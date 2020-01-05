package mdlogger

import (
	"fmt"

	"github.com/short-d/app/fw"
)

var _ fw.Logger = (*Local)(nil)

const datetimeFormat = "2006-01-02 15:04:05"

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

type Local struct {
	prefix         string
	level          fw.LogLevel
	stdout         fw.StdOut
	timer          fw.Timer
	programRuntime fw.ProgramRuntime
}

func (local Local) Fatal(message string) {
	if local.levelAbove(fw.LogFatal) {
		return
	}
	local.log(fw.LogFatalName, message)
}

func (local Local) Error(err error) {
	if local.levelAbove(fw.LogError) {
		return
	}
	local.log(fw.LogErrorName, fmt.Sprintf("%v", err))
}

func (local Local) Warn(message string) {
	if local.levelAbove(fw.LogWarn) {
		return
	}
	local.log(fw.LogWarnName, message)
}

func (local Local) Info(message string) {
	if local.levelAbove(fw.LogInfo) {
		return
	}
	local.log(fw.LogInfoName, message)
}

func (local Local) Debug(message string) {
	if local.levelAbove(fw.LogDebug) {
		return
	}
	local.log(fw.LogDebugName, message)
}

func (local Local) Trace(message string) {
	if local.levelAbove(fw.LogTrace) {
		return
	}
	local.log(fw.LogTraceName, message)
}

func (local Local) log(level fw.LogLevelName, message string) {
	now := local.now()
	caller, err := local.programRuntime.Caller(2)
	if err != nil {
		_, _ = fmt.Fprintf(
			local.stdout,
			"[%s] [%s] %s %s\n",
			local.prefix,
			level,
			now,
			message,
		)
		return
	}
	_, _ = fmt.Fprintf(
		local.stdout,
		"[%s] [%s] %s line %d at %s %s\n",
		local.prefix,
		level,
		now,
		caller.LineNumber,
		caller.FullFilename,
		message,
	)
}

func (local Local) now() string {
	return local.timer.Now().UTC().Format(datetimeFormat)
}

func (local Local) levelAbove(logLevel fw.LogLevel) bool {
	return priorities[local.level] < priorities[logLevel]
}

func NewLocal(
	prefix string,
	level fw.LogLevel,
	stdout fw.StdOut,
	timer fw.Timer,
	programRuntime fw.ProgramRuntime,
) Local {
	return Local{
		prefix:         prefix,
		level:          level,
		stdout:         stdout,
		timer:          timer,
		programRuntime: programRuntime,
	}
}
