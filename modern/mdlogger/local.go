package mdlogger

import (
	"fmt"
	"time"

	"github.com/short-d/app/fw"
)

var _ EntryRepository = (*Local)(nil)

const datetimeFormat = "2006-01-02 15:04:05"

type logLevelName = string

const (
	logFatalName logLevelName = "Fatal"
	logErrorName logLevelName = "Error"
	logWarnName  logLevelName = "Warn"
	logInfoName  logLevelName = "Info"
	logDebugName logLevelName = "Debug"
	logTraceName logLevelName = "Trace"
)

type Local struct {
	stdout fw.StdOut
}

func (local Local) createLogEntry(
	level fw.LogLevel,
	prefix string,
	line int,
	filename string,
	message string,
	date time.Time,
) {
	timeStr := date.Format(datetimeFormat)
	logLevelName := local.getLogLevelName(level)
	message = fmt.Sprintf("line %d at %s %s", line, filename, message)
	_, _ = fmt.Fprintf(
		local.stdout,
		"[%s] [%s] %s %s\n",
		prefix,
		logLevelName,
		timeStr,
		message,
	)
}

func (local Local) getLogLevelName(level fw.LogLevel) string {
	switch level {
	case fw.LogFatal:
		return logFatalName
	case fw.LogError:
		return logErrorName
	case fw.LogWarn:
		return logWarnName
	case fw.LogInfo:
		return logInfoName
	case fw.LogDebug:
		return logDebugName
	case fw.LogTrace:
		return logTraceName
	default:
		return "Should not happen"
	}
}

func NewLocal(stdout fw.StdOut) Local {
	return Local{
		stdout: stdout,
	}
}
