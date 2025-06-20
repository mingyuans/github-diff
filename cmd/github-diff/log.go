package main

import (
	"go.uber.org/zap"
	"sync"
)

var (
	gLogger *zap.Logger
	once    sync.Once
)

func Init(level string) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = ""
	cfg.EncoderConfig.CallerKey = ""
	zapLevel, parseErr := zap.ParseAtomicLevel(level)
	if parseErr != nil {
		panic("Failed to parse log level: " + parseErr.Error())
	}
	cfg.Level = zapLevel
	gLogger, _ = cfg.Build()
}

func Get() *zap.Logger {
	return gLogger
}
