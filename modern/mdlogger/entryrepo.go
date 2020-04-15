package mdlogger

import (
	"time"

	"github.com/short-d/app/fw"
)

type EntryRepository interface {
	createLogEntry(level fw.LogLevel, prefix string, message string, date time.Time)
}
