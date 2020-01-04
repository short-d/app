package mdtest

import "github.com/byliuyang/app/fw"

var _ fw.Logger = (*LoggerFake)(nil)

type LoggerFake struct {
	FatalMessages []string
	Errors        []error
	WarnMessages  []string
	InfoMessages  []string
	DebugMessages []string
	TraceMessages []string
}

func (l *LoggerFake) Fatal(message string) {
	l.FatalMessages = append(l.FatalMessages, message)
}

func (l *LoggerFake) Error(err error) {
	l.Errors = append(l.Errors, err)
}

func (l *LoggerFake) Warn(message string) {
	l.WarnMessages = append(l.WarnMessages, message)
}

func (l *LoggerFake) Info(message string) {
	l.InfoMessages = append(l.InfoMessages, message)
}

func (l *LoggerFake) Debug(message string) {
	l.DebugMessages = append(l.DebugMessages, message)
}

func (l *LoggerFake) Trace(message string) {
	l.TraceMessages = append(l.TraceMessages, message)
}

func NewLoggerFake(
	fatalMessages []string,
	errors []error,
	warnMessages []string,
	infoMessages []string,
	debugMessages []string,
	traceMessages []string,
) LoggerFake {
	return LoggerFake{
		FatalMessages: fatalMessages,
		Errors:        errors,
		WarnMessages:  warnMessages,
		InfoMessages:  infoMessages,
		DebugMessages: debugMessages,
		TraceMessages: traceMessages,
	}
}
