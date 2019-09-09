package mdtest

import "github.com/byliuyang/app/fw"

type logger struct{}

func (logger) Info(info string) {}

func (logger) Error(err error) {}

func (logger) Crash(err error) {}

var LoggerFake fw.Logger = logger{}
