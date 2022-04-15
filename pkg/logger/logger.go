package logger

type LogType int

const (
	DebugLogType LogType = iota
	InfoLogType
	WarningLogType
	ErrorLogType
	FatalLogType
)

var L Logger

type Logger interface {
	Fatal(args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}
