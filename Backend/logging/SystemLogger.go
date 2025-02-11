package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var systemLogger *zap.Logger

// Logger configs for file and console remain the same
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

// Auxiliary functions remain the same
var severityColors = map[zapcore.Level]string{
	zapcore.DebugLevel: "\x1b[36m",
	zapcore.InfoLevel:  "\x1b[32m",
	zapcore.WarnLevel:  "\x1b[33m",
	zapcore.ErrorLevel: "\x1b[31m",
	zapcore.FatalLevel: "\x1b[35m",
}

// Modified to return regular logger instead of sugared
func createLoggerWithCombinedCore(logFile *os.File) *zap.Logger {
	fileLevelEncoder, FileTimeEncoder := createCustomEncoders(false)
	fileEncoderConfig := createBaseEncoderConfig(fileLevelEncoder, FileTimeEncoder)
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(fileEncoderConfig),
		zapcore.AddSync(logFile),
		zap.InfoLevel,
	)

	consoleLevelEncoder, consoleTimeEncoder := createCustomEncoders(true)
	consoleEncoderConfig := createBaseEncoderConfig(consoleLevelEncoder, consoleTimeEncoder)

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	combinedCore := zapcore.NewTee(fileCore, consoleCore)
	return zap.New(combinedCore)
}

// Modified logging functions to use standard logger
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
