package logging

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap/zapcore"
)

func init() {
	setupLogger()
}

func setupLogger() {
	os.Remove("server.log")
	logFile := openLogFile()
	systemLogger = createLoggerWithCombinedCore(logFile)
	requestLogger = createRequestLogger()
}

func openLogFile() *os.File {
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("failed to open log file: " + err.Error())
	}
	return logFile
}

func CloseLogFile() {
	systemLogger.Sync()
}

func createCustomEncoders(isColored bool) (zapcore.LevelEncoder, zapcore.TimeEncoder) {
	levelEncoder := func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		var color string
		if isColored {
			color = severityColors[l]
		} else {
			color = ""
		}
		enc.AppendString(fmt.Sprintf("[ %s%s\x1b[0m ]", color, l.CapitalString()))
	}

	timeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		var timeStr string
		if isColored {
			timeStr = "[ \x1b[90m" + t.Format("2006-01-02 15:04:05") + "\x1b[0m ]"
		} else {
			timeStr = "[ " + t.Format("2006-01-02 15:04:05") + " ]"
		}
		enc.AppendString(timeStr)
	}

	return levelEncoder, timeEncoder
}
