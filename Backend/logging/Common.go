package logging

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	level := zapcore.DebugLevel

	os.Remove("server.log")
	logFile := openLogFile()
	systemLogger = createLoggerWithCombinedCore(logFile, level)
	requestLogger = createRequestLogger()
	SystemWarnLog("Wiping the previous log file")
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
		if isColored {
			enc.AppendString(fmt.Sprintf("[ %s%s\x1b[0m ]", LevelColor(l.CapitalString()), l.CapitalString()))
		} else {
			enc.AppendString(fmt.Sprintf("[ %s ]", l.CapitalString()))
		}
	}
	timeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		if isColored {
			enc.AppendString(fmt.Sprintf("[ \x1b[90m%v\x1b[0m ]", t.Format("2006-01-02 15:04:05")))
		} else {
			enc.AppendString(fmt.Sprintf("[ %v ]", t.Format("2006-01-02 15:04:05")))
		}
	}
	return levelEncoder, timeEncoder
}
