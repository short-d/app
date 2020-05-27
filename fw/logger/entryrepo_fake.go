package logger

import "time"

type Entry struct {
	Level    LogLevel
	Prefix   string
	Line     int
	Filename string
	Message  string
	Date     string
}

var _ EntryRepository = (*EntryRepoFake)(nil)

type EntryRepoFake struct {
	entries []Entry
}

func (f *EntryRepoFake) createLogEntry(
	level LogLevel,
	prefix string,
	line int,
	filename string,
	message string,
	date time.Time,
) {
	f.entries = append(f.entries, Entry{
		Level:    level,
		Prefix:   prefix,
		Line:     line,
		Filename: filename,
		Message:  message,
		Date:     date.Format(time.RFC3339),
	})
}

func (f EntryRepoFake) GetEntries() []Entry {
	return f.entries
}

func NewEntryRepoFake() EntryRepoFake {
	return EntryRepoFake{
		entries: make([]Entry, 0),
	}
}
