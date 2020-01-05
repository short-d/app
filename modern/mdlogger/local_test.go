package mdlogger

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/short-d/app/fw"
	"github.com/short-d/app/mdtest"
)

func TestLocal_Fatal(t *testing.T) {
	testCases := []struct {
		name           string
		logLevel       fw.LogLevel
		now            string
		prefix         string
		messages       []string
		callers        []fw.Caller
		expectedOutput string
	}{
		{
			name:           "logging disabled",
			logLevel:       fw.LogOff,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{},
			expectedOutput: "",
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
			expectedOutput: `[Prefix] [Fatal] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 1
[Prefix] [Fatal] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 2
`,
		},
		{
			name:     "log error message",
			logLevel: fw.LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Fatal] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
		{
			name:     "log warn message",
			logLevel: fw.LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Fatal] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
		{
			name:     "log info message",
			logLevel: fw.LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Fatal] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
		{
			name:     "log debug message",
			logLevel: fw.LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Fatal] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Fatal] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			var buf bytes.Buffer
			logger := NewLocal(
				testCase.prefix,
				testCase.logLevel,
				&buf,
				timer,
				programRuntime,
			)

			for _, message := range testCase.messages {
				logger.Fatal(message)
			}
			gotOutput := buf.String()
			mdtest.Equal(t, testCase.expectedOutput, gotOutput)
		})
	}
}

func TestLocal_Error(t *testing.T) {
	testCases := []struct {
		name           string
		logLevel       fw.LogLevel
		now            string
		prefix         string
		messages       []string
		callers        []fw.Caller
		expectedOutput string
	}{
		{
			name:           "logging disabled",
			logLevel:       fw.LogOff,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{},
			expectedOutput: "",
		},
		{
			name:           "log fatal message",
			logLevel:       fw.LogFatal,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{},
			callers:        []fw.Caller{},
			expectedOutput: "",
		},
		{
			name:     "log error message",
			logLevel: fw.LogError,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Error] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 1
[Prefix] [Error] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 2
`,
		},
		{
			name:     "log warn message",
			logLevel: fw.LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Error] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
		{
			name:     "log info message",
			logLevel: fw.LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Error] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
		{
			name:     "log debug message",
			logLevel: fw.LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Error] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Error] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			var buf bytes.Buffer
			logger := NewLocal(
				testCase.prefix,
				testCase.logLevel,
				&buf,
				timer,
				programRuntime,
			)

			for _, message := range testCase.messages {
				logger.Error(errors.New(message))
			}
			gotOutput := buf.String()
			mdtest.Equal(t, testCase.expectedOutput, gotOutput)
		})
	}
}

func TestLocal_Warn(t *testing.T) {
	testCases := []struct {
		name           string
		logLevel       fw.LogLevel
		now            string
		prefix         string
		messages       []string
		callers        []fw.Caller
		expectedOutput string
	}{
		{
			name:           "logging disabled",
			logLevel:       fw.LogOff,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{},
			expectedOutput: "",
		},
		{
			name:           "log fatal message",
			logLevel:       fw.LogFatal,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{},
			callers:        []fw.Caller{},
			expectedOutput: "",
		},
		{
			name:           "log error message",
			logLevel:       fw.LogError,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: "",
		},
		{
			name:     "log warn message",
			logLevel: fw.LogWarn,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Warn] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 1
[Prefix] [Warn] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 2
`,
		},
		{
			name:     "log info message",
			logLevel: fw.LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Warn] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
		{
			name:     "log debug message",
			logLevel: fw.LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Warn] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Warn] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			var buf bytes.Buffer
			logger := NewLocal(
				testCase.prefix,
				testCase.logLevel,
				&buf,
				timer,
				programRuntime,
			)

			for _, message := range testCase.messages {
				logger.Warn(message)
			}
			gotOutput := buf.String()
			mdtest.Equal(t, testCase.expectedOutput, gotOutput)
		})
	}
}

func TestLocal_Info(t *testing.T) {
	testCases := []struct {
		name           string
		logLevel       fw.LogLevel
		now            string
		prefix         string
		messages       []string
		callers        []fw.Caller
		expectedOutput string
	}{
		{
			name:           "logging disabled",
			logLevel:       fw.LogOff,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{},
			expectedOutput: "",
		},
		{
			name:           "log fatal message",
			logLevel:       fw.LogFatal,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{},
			callers:        []fw.Caller{},
			expectedOutput: "",
		},
		{
			name:           "log error message",
			logLevel:       fw.LogError,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: "",
		},
		{
			name:           "log warn message",
			logLevel:       fw.LogWarn,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: "",
		},
		{
			name:     "log info message",
			logLevel: fw.LogInfo,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Info] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 1
[Prefix] [Info] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 2
`,
		},
		{
			name:     "log debug message",
			logLevel: fw.LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Info] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Info] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			var buf bytes.Buffer
			logger := NewLocal(
				testCase.prefix,
				testCase.logLevel,
				&buf,
				timer,
				programRuntime,
			)

			for _, message := range testCase.messages {
				logger.Info(message)
			}
			gotOutput := buf.String()
			mdtest.Equal(t, testCase.expectedOutput, gotOutput)
		})
	}
}

func TestLocal_Debug(t *testing.T) {
	testCases := []struct {
		name           string
		logLevel       fw.LogLevel
		now            string
		prefix         string
		messages       []string
		callers        []fw.Caller
		expectedOutput string
	}{
		{
			name:           "logging disabled",
			logLevel:       fw.LogOff,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{},
			expectedOutput: "",
		},
		{
			name:           "log fatal message",
			logLevel:       fw.LogFatal,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{},
			callers:        []fw.Caller{},
			expectedOutput: "",
		},
		{
			name:           "log error message",
			logLevel:       fw.LogError,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: "",
		},
		{
			name:           "log warn message",
			logLevel:       fw.LogWarn,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: "",
		},
		{
			name:           "log info message",
			logLevel:       fw.LogInfo,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: "",
		},
		{
			name:     "log debug message",
			logLevel: fw.LogDebug,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Debug] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 1
[Prefix] [Debug] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 2
`,
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Debug] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message
`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			var buf bytes.Buffer
			logger := NewLocal(
				testCase.prefix,
				testCase.logLevel,
				&buf,
				timer,
				programRuntime,
			)

			for _, message := range testCase.messages {
				logger.Debug(message)
			}
			gotOutput := buf.String()
			mdtest.Equal(t, testCase.expectedOutput, gotOutput)
		})
	}
}

func TestLocal_Trace(t *testing.T) {
	testCases := []struct {
		name           string
		logLevel       fw.LogLevel
		now            string
		prefix         string
		messages       []string
		callers        []fw.Caller
		expectedOutput string
	}{
		{
			name:           "logging disabled",
			logLevel:       fw.LogOff,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{},
			expectedOutput: "",
		},
		{
			name:           "log fatal message",
			logLevel:       fw.LogFatal,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{},
			callers:        []fw.Caller{},
			expectedOutput: "",
		},
		{
			name:           "log error message",
			logLevel:       fw.LogError,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: "",
		},
		{
			name:           "log warn message",
			logLevel:       fw.LogWarn,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: "",
		},
		{
			name:           "log info message",
			logLevel:       fw.LogInfo,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: "",
		},
		{
			name:           "log debug message",
			logLevel:       fw.LogDebug,
			now:            "2020-01-04T10:20:04-00:00",
			prefix:         "Prefix",
			messages:       []string{"message"},
			callers:        []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: "",
		},
		{
			name:     "log trace message",
			logLevel: fw.LogTrace,
			now:      "2020-01-04T10:20:04-00:00",
			prefix:   "Prefix",
			messages: []string{"message 1", "message 2"},
			callers:  []fw.Caller{{}, {}, {FullFilename: "github.com/short-d/app/test.go", LineNumber: 10}},
			expectedOutput: `[Prefix] [Trace] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 1
[Prefix] [Trace] 2020-01-04 10:20:04 line 10 at github.com/short-d/app/test.go message 2
`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			now, err := time.Parse(time.RFC3339, testCase.now)
			mdtest.Equal(t, nil, err)
			timer := mdtest.NewTimerFake(now)

			programRuntime, err := mdtest.NewProgramRuntimeFake(testCase.callers)
			mdtest.Equal(t, nil, err)

			var buf bytes.Buffer
			logger := NewLocal(
				testCase.prefix,
				testCase.logLevel,
				&buf,
				timer,
				programRuntime,
			)

			for _, message := range testCase.messages {
				logger.Trace(message)
			}
			gotOutput := buf.String()
			mdtest.Equal(t, testCase.expectedOutput, gotOutput)
		})
	}
}
