package mdtest

import "github.com/byliuyang/app/fw"

var _ fw.Logger = (*LoggerFake)(nil)

type LoggerFake struct {
	Infos   []string
	Errors  []error
	Crashes []error
}

func (l *LoggerFake) Info(info string) {
	l.Infos = append(l.Infos, info)
}

func (l *LoggerFake) Error(err error) {
	l.Errors = append(l.Errors, err)
}

func (l *LoggerFake) Crash(err error) {
	l.Crashes = append(l.Crashes, err)
}

func NewLoggerFake() LoggerFake {
	return LoggerFake{
		Infos:   make([]string, 0),
		Errors:  make([]error, 0),
		Crashes: make([]error, 0),
	}
}
