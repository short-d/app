package mdlogger

import (
	"errors"
	"testing"
	"time"

	"github.com/short-d/app/fw"
	"github.com/short-d/app/mdtest"
)

type entry struct {
	level     fw.LogLevelName
	prefix    string
	message   string
	timestamp string
}

var _ EntryRepository = (*FakeEntryRepo)(nil)

type FakeEntryRepo struct {
	entries []entry
}

func (f *FakeEntryRepo) createLogEntry(level fw.LogLevelName, prefix string, message string, timestamp time.Time) {
	f.entries = append(f.entries, entry{
		level:     level,
		prefix:    prefix,
		message:   message,
		timestamp: timestamp.Format(time.RFC3339),
	})
}

func TestLogger_Fatal(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        fw.LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []fw.Caller
		expectedEntries []entry
	}{
		{
			name:     "logging disabled",
			logLevel: fw.LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{},
		},
		{
			name:     "log fatal message",
			logLevel: fw.LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers: []fw.Caller{{}, {}, {
				FullFilename: "github.com/short-d/app/test.go",
				LineNumber:   10,
			}},
			expectedEntries: []entry{
				{
					level:     fw.LogFatalName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 1",
					timestamp: "2020-01-04T10:20:04Z",
				},
				{
					level:     fw.LogFatalName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 2",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log error message",
			logLevel: fw.LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogFatalName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log warn message",
			logLevel: fw.LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogFatalName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log info message",
			logLevel: fw.LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogFatalName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log debug message",
			logLevel: fw.LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogFatalName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogFatalName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			fakeEntryRepo := &FakeEntryRepo{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				timer,
				programRuntime,
				fakeEntryRepo,
			)

			for _, message := range testCase.messages {
				logger.Fatal(message)
			}

			mdtest.Equal(t, testCase.expectedEntries, fakeEntryRepo.entries)
		})
	}
}

func TestLogger_Error(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        fw.LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []fw.Caller
		expectedEntries []entry
	}{
		{
			name:     "logging disabled",
			logLevel: fw.LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
		},
		{
			name:     "log fatal message",
			logLevel: fw.LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{},
		},
		{
			name:     "log error message",
			logLevel: fw.LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogErrorName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 1",
					timestamp: "2020-01-04T10:20:04Z",
				},
				{
					level:     fw.LogErrorName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 2",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log warn message",
			logLevel: fw.LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogErrorName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log info message",
			logLevel: fw.LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogErrorName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log debug message",
			logLevel: fw.LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogErrorName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogErrorName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			fakeEntryRepo := &FakeEntryRepo{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				timer,
				programRuntime,
				fakeEntryRepo,
			)

			for _, message := range testCase.messages {
				logger.Error(errors.New(message))
			}

			mdtest.Equal(t, testCase.expectedEntries, fakeEntryRepo.entries)
		})
	}
}

func TestLogger_Warn(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        fw.LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []fw.Caller
		expectedEntries []entry
	}{
		{
			name:     "logging disabled",
			logLevel: fw.LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{},
		},
		{
			name:     "log fatal message",
			logLevel: fw.LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{},
			callers:  []fw.Caller{},
		},
		{
			name:     "log error message",
			logLevel: fw.LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log warn message",
			logLevel: fw.LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogWarnName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 1",
					timestamp: "2020-01-04T10:20:04Z",
				},
				{
					level:     fw.LogWarnName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 2",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log info message",
			logLevel: fw.LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogWarnName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log debug message",
			logLevel: fw.LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogWarnName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogWarnName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			fakeEntryRepo := &FakeEntryRepo{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				timer,
				programRuntime,
				fakeEntryRepo,
			)

			for _, message := range testCase.messages {
				logger.Warn(message)
			}

			mdtest.Equal(t, testCase.expectedEntries, fakeEntryRepo.entries)
		})
	}
}

func TestLogger_Info(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        fw.LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []fw.Caller
		expectedEntries []entry
	}{
		{
			name:     "logging disabled",
			logLevel: fw.LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{},
		},
		{
			name:     "log fatal message",
			logLevel: fw.LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{},
			callers:  []fw.Caller{},
		},
		{
			name:     "log error message",
			logLevel: fw.LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log warn message",
			logLevel: fw.LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log info message",
			logLevel: fw.LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogInfoName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 1",
					timestamp: "2020-01-04T10:20:04Z",
				},
				{
					level:     fw.LogInfoName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 2",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log debug message",
			logLevel: fw.LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogInfoName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogInfoName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			fakeEntryRepo := &FakeEntryRepo{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				timer,
				programRuntime,
				fakeEntryRepo,
			)

			for _, message := range testCase.messages {
				logger.Info(message)
			}

			mdtest.Equal(t, testCase.expectedEntries, fakeEntryRepo.entries)
		})
	}
}

func TestLogger_Debug(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        fw.LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []fw.Caller
		expectedEntries []entry
	}{
		{
			name:     "logging disabled",
			logLevel: fw.LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{},
		},
		{
			name:     "log fatal message",
			logLevel: fw.LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{},
			callers:  []fw.Caller{},
		},
		{
			name:     "log error message",
			logLevel: fw.LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log warn message",
			logLevel: fw.LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log info message",
			logLevel: fw.LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log debug message",
			logLevel: fw.LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogDebugName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 1",
					timestamp: "2020-01-04T10:20:04Z",
				},
				{
					level:     fw.LogDebugName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 2",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogDebugName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			fakeEntryRepo := &FakeEntryRepo{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				timer,
				programRuntime,
				fakeEntryRepo,
			)

			for _, message := range testCase.messages {
				logger.Debug(message)
			}

			mdtest.Equal(t, testCase.expectedEntries, fakeEntryRepo.entries)
		})
	}
}

func TestLogger_Trace(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        fw.LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []fw.Caller
		expectedEntries []entry
	}{
		{
			name:     "logging disabled",
			logLevel: fw.LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{},
		},
		{
			name:     "log fatal message",
			logLevel: fw.LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{},
			callers:  []fw.Caller{},
		},
		{
			name:     "log error message",
			logLevel: fw.LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log warn message",
			logLevel: fw.LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log info message",
			logLevel: fw.LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log debug message",
			logLevel: fw.LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []entry{
				{
					level:     fw.LogTraceName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 1",
					timestamp: "2020-01-04T10:20:04Z",
				},
				{
					level:     fw.LogTraceName,
					prefix:    "Prefix",
					message:   "line 10 at github.com/short-d/app/test.go message 2",
					timestamp: "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			fakeEntryRepo := &FakeEntryRepo{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				timer,
				programRuntime,
				fakeEntryRepo,
			)

			for _, message := range testCase.messages {
				logger.Trace(message)
			}

			mdtest.Equal(t, testCase.expectedEntries, fakeEntryRepo.entries)
		})
	}
}
