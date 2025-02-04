package utils

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}

	encoderConfigLogFile := zapcore.EncoderConfig{
		TimeKey:    "timestamp",
		LevelKey:   "level",
		MessageKey: "msg",
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(fmt.Sprintf("[ %v ]", l.CapitalString()))
		},
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(fmt.Sprintf("[ %v ]", t.Format("2006-01-02 15:04:05")))
		},
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " ",
	}

	// Create two cores - one for console, one for file
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfigLogFile),
		zapcore.AddSync(logFile),
		zap.InfoLevel,
	)

	encoderConfigConsole := zapcore.EncoderConfig{
		TimeKey:    "timestamp",
		LevelKey:   "level",
		MessageKey: "msg",
		EncodeLevel: func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			switch l {
			case zapcore.DebugLevel:
				enc.AppendString("[ \x1b[36m" + l.CapitalString() + "\x1b[0m ]") // Cyan
			case zapcore.InfoLevel:
				enc.AppendString("[ \x1b[32m" + l.CapitalString() + "\x1b[0m ]") // Green
			case zapcore.WarnLevel:
				enc.AppendString("[ \x1b[33m" + l.CapitalString() + "\x1b[0m ]") // Yellow
			case zapcore.ErrorLevel:
				enc.AppendString("[ \x1b[31m" + l.CapitalString() + "\x1b[0m ]") // Red
			case zapcore.FatalLevel:
				enc.AppendString("[ \x1b[35m" + l.CapitalString() + "\x1b[0m ]") // Magenta
			default:
				enc.AppendString("[ " + l.CapitalString() + " ]")
			}
		},
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString("[ \x1b[90m" + t.Format("2006-01-02 15:04:05") + "\x1b[0m ]") // Gray
		},
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " ",
	}

	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfigConsole),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	// Combine both cores using zapcore.NewTee
	core := zapcore.NewTee(consoleCore, fileCore)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)).Sugar()
}
