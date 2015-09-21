package kcl

var logger Logger

type Logger interface {
	Log(msg string)
	Init()
}

func init() {
	SetLogger(&nullLogger{})
}

func SetLogger(l Logger) {
	logger = l
	logger.Init()
}

type nullLogger struct {
}

func (l *nullLogger) Init() {
}

func (l *nullLogger) Log(msg string) {
}
