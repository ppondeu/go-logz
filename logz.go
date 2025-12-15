package logz

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger


func InitLog(serverMode string) {
	var config zap.Config
	
	if serverMode == "development" {
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config = zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
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

func Fatalf(format string, args ...interface{}) {
	log.Sugar().Fatalf(format, args...)
}

func Warn(message string, fields ...zapcore.Field) {
	log.Warn(message, fields...)
}

func Errorf(format string, args ...interface{}) {
	log.Sugar().Errorf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Sugar().Infof(format, args...)
}