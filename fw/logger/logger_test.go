package logger

import (
	"errors"
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/runtime"
	"github.com/short-d/app/fw/timer"
)

func TestLogger_Fatal(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []runtime.Caller
		expectedEntries []Entry
	}{
		{
			name:     "logging disabled",
			logLevel: LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{},
		},
		{
			name:     "log fatal message",
			logLevel: LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogFatal,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 1",
					Date:     "2020-01-04T10:20:04Z",
				},
				{
					Level:    LogFatal,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 2",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log error message",
			logLevel: LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogFatal,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log warn message",
			logLevel: LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogFatal,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log info message",
			logLevel: LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogFatal,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log debug message",
			logLevel: LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogFatal,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log trace message",
			logLevel: LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogFatal,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			assert.Equal(t, nil, err)
			tm := timer.NewStub(now)

			programRuntime, err := runtime.NewFake(testCase.callers)
			assert.Equal(t, nil, err)

			EntryRepoFake := &EntryRepoFake{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				tm,
				programRuntime,
				EntryRepoFake,
			)

			for _, message := range testCase.messages {
				logger.Fatal(message)
			}

			assert.Equal(t, testCase.expectedEntries, EntryRepoFake.entries)
		})
	}
}

func TestLogger_Error(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []runtime.Caller
		expectedEntries []Entry
	}{
		{
			name:     "logging disabled",
			logLevel: LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
		},
		{
			name:     "log fatal message",
			logLevel: LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{},
		},
		{
			name:     "log error message",
			logLevel: LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogError,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 1",
					Date:     "2020-01-04T10:20:04Z",
				},
				{
					Level:    LogError,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 2",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log warn message",
			logLevel: LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogError,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log info message",
			logLevel: LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogError,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log debug message",
			logLevel: LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogError,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log trace message",
			logLevel: LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogError,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			assert.Equal(t, nil, err)
			tm := timer.NewStub(now)

			programRuntime, err := runtime.NewFake(testCase.callers)
			assert.Equal(t, nil, err)

			EntryRepoFake := &EntryRepoFake{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				tm,
				programRuntime,
				EntryRepoFake,
			)

			for _, message := range testCase.messages {
				logger.Error(errors.New(message))
			}

			assert.Equal(t, testCase.expectedEntries, EntryRepoFake.entries)
		})
	}
}

func TestLogger_Warn(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []runtime.Caller
		expectedEntries []Entry
	}{
		{
			name:     "logging disabled",
			logLevel: LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{},
		},
		{
			name:     "log fatal message",
			logLevel: LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{},
			callers:  []runtime.Caller{},
		},
		{
			name:     "log error message",
			logLevel: LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log warn message",
			logLevel: LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogWarn,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 1",
					Date:     "2020-01-04T10:20:04Z",
				},
				{
					Level:    LogWarn,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 2",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log info message",
			logLevel: LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogWarn,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log debug message",
			logLevel: LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogWarn,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log trace message",
			logLevel: LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogWarn,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			assert.Equal(t, nil, err)
			tm := timer.NewStub(now)

			programRuntime, err := runtime.NewFake(testCase.callers)
			assert.Equal(t, nil, err)

			EntryRepoFake := &EntryRepoFake{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				tm,
				programRuntime,
				EntryRepoFake,
			)

			for _, message := range testCase.messages {
				logger.Warn(message)
			}

			assert.Equal(t, testCase.expectedEntries, EntryRepoFake.entries)
		})
	}
}

func TestLogger_Info(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []runtime.Caller
		expectedEntries []Entry
	}{
		{
			name:     "logging disabled",
			logLevel: LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{},
		},
		{
			name:     "log fatal message",
			logLevel: LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{},
			callers:  []runtime.Caller{},
		},
		{
			name:     "log error message",
			logLevel: LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log warn message",
			logLevel: LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log info message",
			logLevel: LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogInfo,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 1",
					Date:     "2020-01-04T10:20:04Z",
				},
				{
					Level:    LogInfo,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 2",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log debug message",
			logLevel: LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogInfo,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log trace message",
			logLevel: LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogInfo,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			assert.Equal(t, nil, err)
			tm := timer.NewStub(now)

			programRuntime, err := runtime.NewFake(testCase.callers)
			assert.Equal(t, nil, err)

			EntryRepoFake := &EntryRepoFake{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				tm,
				programRuntime,
				EntryRepoFake,
			)

			for _, message := range testCase.messages {
				logger.Info(message)
			}

			assert.Equal(t, testCase.expectedEntries, EntryRepoFake.entries)
		})
	}
}

func TestLogger_Debug(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []runtime.Caller
		expectedEntries []Entry
	}{
		{
			name:     "logging disabled",
			logLevel: LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{},
		},
		{
			name:     "log fatal message",
			logLevel: LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{},
			callers:  []runtime.Caller{},
		},
		{
			name:     "log error message",
			logLevel: LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log warn message",
			logLevel: LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log info message",
			logLevel: LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log debug message",
			logLevel: LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogDebug,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 1",
					Date:     "2020-01-04T10:20:04Z",
				},
				{
					Level:    LogDebug,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 2",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
		{
			name:     "log trace message",
			logLevel: LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogDebug,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			assert.Equal(t, nil, err)
			tm := timer.NewStub(now)

			programRuntime, err := runtime.NewFake(testCase.callers)
			assert.Equal(t, nil, err)

			EntryRepoFake := &EntryRepoFake{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				tm,
				programRuntime,
				EntryRepoFake,
			)

			for _, message := range testCase.messages {
				logger.Debug(message)
			}

			assert.Equal(t, testCase.expectedEntries, EntryRepoFake.entries)
		})
	}
}

func TestLogger_Trace(t *testing.T) {
	testCases := []struct {
		name            string
		logLevel        LogLevel
		now             string
		prefix          string
		messages        []string
		callers         []runtime.Caller
		expectedEntries []Entry
	}{
		{
			name:     "logging disabled",
			logLevel: LogOff,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{},
		},
		{
			name:     "log fatal message",
			logLevel: LogFatal,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{},
			callers:  []runtime.Caller{},
		},
		{
			name:     "log error message",
			logLevel: LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log warn message",
			logLevel: LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log info message",
			logLevel: LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log debug message",
			logLevel: LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
		},
		{
			name:     "log trace message",
			logLevel: LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []runtime.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedEntries: []Entry{
				{
					Level:    LogTrace,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 1",
					Date:     "2020-01-04T10:20:04Z",
				},
				{
					Level:    LogTrace,
					Prefix:   "Prefix",
					Line:     10,
					Filename: "github.com/short-d/app/test.go",
					Message:  "message 2",
					Date:     "2020-01-04T10:20:04Z",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			assert.Equal(t, nil, err)
			tm := timer.NewStub(now)

			programRuntime, err := runtime.NewFake(testCase.callers)
			assert.Equal(t, nil, err)

			EntryRepoFake := &EntryRepoFake{}

			logger := NewLogger(
				testCase.prefix,
				testCase.logLevel,
				tm,
				programRuntime,
				EntryRepoFake,
			)

			for _, message := range testCase.messages {
				logger.Trace(message)
			}

			assert.Equal(t, testCase.expectedEntries, EntryRepoFake.entries)
		})
	}
}
