package logger

import (
	"time"
)

type EntryRepository interface {
	createLogEntry(
		level LogLevel,
		prefix string,
		line int,
		filename string,
		message string,
		date time.Time,
	)
}
