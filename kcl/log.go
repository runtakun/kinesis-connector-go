package kcl

var logger Logger

type Logger interface {
	Printf(format string, v ...interface{})
}

func init() {
	SetLogger(&nullLogger{})
}

func SetLogger(l Logger) {
	logger = l
}

type nullLogger struct {
}

func (l *nullLogger) Printf(format string, v ...interface{}) {
}
