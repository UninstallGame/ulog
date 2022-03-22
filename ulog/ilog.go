package ulog

type ILog interface {
	SetLogLevel(logLevel int)
	Debug(text string)
	Info(text string)
	Warning(text string)
	Error(text string, err error)
	Fatal(text string, err error)
}
