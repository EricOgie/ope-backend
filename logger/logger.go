package logger

import (
	"github.com/EricOgie/ope-be/konstants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var appLog *zap.Logger

// Zap Logger initiation
func init() {
	var err error
	myConfig := zap.NewProductionConfig()
	myEncoderConfig := zap.NewProductionEncoderConfig()

	// Setting myEncoderConfig details
	myEncoderConfig.TimeKey = konstants.TIME
	myEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	myEncoderConfig.LevelKey = konstants.LOGGER_TYPE
	myEncoderConfig.StacktraceKey = "" // This will make the logger omit stacktrace

	// Next, we set the encoderconfiguration attribute of the myConfig to our custom EncoderConfig
	myConfig.EncoderConfig = myEncoderConfig

	appLog, err = myConfig.Build(zap.AddCallerSkip(1))
	// NB: The zap.AddCallerSkip(1) will enable the logger config to register file and line of call
	if err != nil {
		panic(err)
	}
}

// Info helper func to expose the Info logging
func Info(msg string, fields ...zap.Field) {
	appLog.Info(msg, fields...)
}

// Debug helper func to expose the Info logging
func Debug(msg string, fields ...zap.Field) {
	appLog.Debug(msg, fields...)
}

// Error  helper func to expose the Info logging
func Error(msg string, fields ...zap.Field) {
	appLog.Error(msg, fields...)
}
