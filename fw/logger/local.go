package logger

import (
	"fmt"
	"time"

	"github.com/short-d/app/fw/io"
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
	output      io.Output
	showLogLine bool
}

func (l Local) createLogEntry(
	level LogLevel,
	prefix string,
	line int,
	filename string,
	message string,
	date time.Time,
) {
	timeStr := date.Format(datetimeFormat)
	logLevelName := l.getLogLevelName(level)
	message = l.message(line, filename, message)
	_, _ = fmt.Fprintf(
		l.output,
		"[%s] [%s] %s %s\n",
		prefix,
		logLevelName,
		timeStr,
		message,
	)
}

func (l Local) message(line int,
	filename string,
	message string) string {
	if !l.showLogLine {
		return message
	}
	return fmt.Sprintf("line %d at %s %s", line, filename, message)
}

func (l Local) getLogLevelName(level LogLevel) string {
	switch level {
	case LogFatal:
		return logFatalName
	case LogError:
		return logErrorName
	case LogWarn:
		return logWarnName
	case LogInfo:
		return logInfoName
	case LogDebug:
		return logDebugName
	case LogTrace:
		return logTraceName
	default:
		return "Should not happen"
	}
}

func NewLocal(output io.Output, showLogLine bool) Local {
	return Local{
		output:      output,
		showLogLine: showLogLine,
	}
}
