package logz

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log   *zap.Logger
	mutex sync.Mutex
)

func initializeDefault() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return l
}

func InitLog(serverMode string, opts ...func(*zap.Config)) {
	mutex.Lock()
	defer mutex.Unlock()

	var cfg zap.Config

	if serverMode == "development" {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig.TimeKey = "timestamp"
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	l, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	log = l
}

func getLogger() *zap.Logger {
	mutex.Lock()
	defer mutex.Unlock()

	if log == nil {
		log = initializeDefault()
	}
	return log
}

func Info(message string, fields ...zapcore.Field) {
	getLogger().Info(message, fields...)
}

func Debug(message string, fields ...zapcore.Field) {
	getLogger().Debug(message, fields...)
}

func Warn(message string, fields ...zapcore.Field) {
	getLogger().Warn(message, fields...)
}

func Error(message any, fields ...zapcore.Field) {
	switch v := message.(type) {
	case error:
		getLogger().Error(v.Error(), fields...)
	case string:
		getLogger().Error(v, fields...)
	}
}

func Fatal(message any, fields ...zapcore.Field) {
	switch v := message.(type) {
	case error:
		getLogger().Fatal(v.Error(), fields...)
	case string:
		getLogger().Fatal(v, fields...)
	}
}

func Infof(format string, args ...interface{}) {
	getLogger().Sugar().Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	getLogger().Sugar().Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	getLogger().Sugar().Fatalf(format, args...)
}
