package mdlogger

import (
	"fmt"
	"runtime"

	"github.com/byliuyang/app/fw"
)

var _ fw.Logger = (*Local)(nil)

const datetimeFormat  = "2006-01-02 15:04:05"

type Local struct {
	prefix string
	timer  fw.Timer
	level  fw.LogLevel
}

func (local Local) Fatal(message string) {
	if local.levelAboveFatal() {
		return
	}
	local.log(fw.LogFatalName, message)
}

func (local Local) Error(err error) {
	if local.levelAboveError() {
		return
	}
	local.log(fw.LogFatalName, fmt.Sprintf("%v", err))
}

func (local Local) Warn(message string) {
	if local.levelAboveWarn() {
		return
	}
	local.log(fw.LogWarnName, message)
}

func (local Local) Info(message string) {
	if local.levelAboveInfo() {
		return
	}
	local.log(fw.LogInfoName, message)
}

func (local Local) Debug(message string) {
	if local.levelAboveDebug() {
		return
	}
	local.log(fw.LogDebugName, message)
}

func (local Local) Trace(message string) {
	if local.levelAboveTrace() {
		return
	}
	local.log(fw.LogTraceName, message)
}

func (local Local) log(level fw.LogLevelName, message string) {
	now := local.timer.Now().UTC().Format(datetimeFormat)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Printf("[%s] [%s] %s %s\n", local.prefix, level, now, message)
		return
	}
	fmt.Printf(
		"[%s] [%s] %s line %d at %s %s\n",
		local.prefix,
		level,
		now,
		line,
		file,
		message,
	)
}

func (local Local) levelAboveFatal() bool {
	return local.level == fw.LogOff
}

func (local Local) levelAboveError() bool {
	if local.levelAboveFatal() {
		return true
	}
	return local.level == fw.LogFatal
}

func (local Local) levelAboveWarn() bool {
	if local.levelAboveError() {
		return true
	}
	return local.level == fw.LogError
}

func (local Local) levelAboveInfo() bool {
	if local.levelAboveWarn() {
		return true
	}
	return local.level == fw.LogWarn
}

func (local Local) levelAboveDebug() bool {
	if local.levelAboveInfo() {
		return true
	}
	return local.level == fw.LogInfo
}

func (local Local) levelAboveTrace() bool {
	if local.levelAboveDebug() {
		return true
	}
	return local.level == fw.LogDebug
}

func NewLocal(prefix string, timer fw.Timer, level fw.LogLevel) Local {
	return Local{
		prefix: prefix,
		timer:  timer,
		level:  level,
	}
}
