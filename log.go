package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetLogger(logLevel string) *zap.Logger {
	var config zap.Config
	config = zap.NewProductionConfig()
	if logLevel == "debug" {
		config.Level.SetLevel(zapcore.DebugLevel)
	}
	if logLevel == "info" {
		config.Level.SetLevel(zapcore.InfoLevel)
	}
	if logLevel == "warn" {
		config.Level.SetLevel(zapcore.WarnLevel)
	}
	if logLevel == "error" {
		config.Level.SetLevel(zapcore.ErrorLevel)
	}
	config.Encoding = "console"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//config.DisableCaller = true
	config.DisableStacktrace = true
	config.OutputPaths = []string{"stdout"}
	//config.OutputPaths = []string{"server.log"}
	logger, err := config.Build()
	defer logger.Sync()
	if err != nil {
		logger.Error("SetLogger", zap.Error(err))
	}
	zap.ReplaceGlobals(logger)
	return logger
}
