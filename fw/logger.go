package fw

type LogLevel int

const (
	LogFatal LogLevel = iota
	LogError
	LogWarn
	LogInfo
	LogDebug
	LogTrace
	LogOff
)

type Logger interface {
	Fatal(message string)
	Error(err error)
	Warn(message string)
	Info(message string)
	Debug(message string)
	Trace(message string)
}
