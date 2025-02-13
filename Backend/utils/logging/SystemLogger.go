package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var systemLogger *zap.Logger

func createBaseEncoderConfig(levelEncoder zapcore.LevelEncoder, timeEncoder zapcore.TimeEncoder) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:          "timestamp",
		LevelKey:         "level",
		MessageKey:       "msg",
		EncodeLevel:      levelEncoder,
		EncodeTime:       timeEncoder,
		ConsoleSeparator: " ",
	}
}

// Modified to return regular logger instead of sugared
func createLoggerWithCombinedCore(logFile *os.File, level zapcore.Level) *zap.Logger {
	fileLevelEncoder, FileTimeEncoder := createCustomEncoders(false)
	fileEncoderConfig := createBaseEncoderConfig(fileLevelEncoder, FileTimeEncoder)
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(fileEncoderConfig),
		zapcore.AddSync(logFile),
		level,
	)

	consoleLevelEncoder, consoleTimeEncoder := createCustomEncoders(true)
	consoleEncoderConfig := createBaseEncoderConfig(consoleLevelEncoder, consoleTimeEncoder)

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoderConfig),
		zapcore.AddSync(os.Stdout),
		level,
	)

	combinedCore := zapcore.NewTee(fileCore, consoleCore)
	return zap.New(combinedCore)
}

func SystemInfoLog(msg string) {
	systemLogger.Info(msg)
}

func SystemDebugLog(msg string) {
	systemLogger.Debug(msg)
}

func SystemWarnLog(msg string) {
	systemLogger.Warn(msg)
}

func SystemErrorLog(msg string) {
	systemLogger.Error(msg)
}

func SystemFatalLog(msg string) {
	systemLogger.Fatal(msg)
	panic("")
}
