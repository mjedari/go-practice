package services

import "github.com/sirupsen/logrus"

type Logger struct {
	//
	*logrus.Logger
}

func NewLogger() *Logger {
	logger := logrus.New()
	return &Logger{Logger: logger}
}

func (l Logger) Decorate(function IDecoratedFunction) IDecoratedFunction {
	l.Info("you hit the log decorator")
	//
	return function
}
