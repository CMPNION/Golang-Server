package logger

import "go.uber.org/zap"

func NewLogger() *zap.Logger {

	l, _ := zap.NewDevelopment()
	l.Sugar().Info("Logger started")

	return l
}
