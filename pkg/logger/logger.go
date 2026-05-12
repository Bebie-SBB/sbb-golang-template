package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

var log *Logger

func init() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLog, _ := cfg.Build()
	log = &Logger{zapLog.Sugar()}
}

func Get() *Logger {
	return log
}
