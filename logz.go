package logz

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

type Option func(*zap.Config)

func InitLog(serverMode string, opts ...Option) {
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

	var err error
	log, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zapcore.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zapcore.Field) {
	log.Debug(message, fields...)
}

func Warn(message string, fields ...zapcore.Field) {
	log.Warn(message, fields...)
}

func Error(message any, fields ...zapcore.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}

func Fatal(message any, fields ...zapcore.Field) {
	switch v := message.(type) {
	case error:
		log.Fatal(v.Error(), fields...)
	case string:
		log.Fatal(v, fields...)
	}
}

func Infof(format string, args ...interface{}) {
	log.Sugar().Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Sugar().Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Sugar().Fatalf(format, args...)
}
