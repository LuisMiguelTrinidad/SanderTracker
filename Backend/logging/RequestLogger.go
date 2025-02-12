package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var requestLogger *zap.Logger

func formatRequestMessage(msg, method string, status int) string {
	resetColor := "\x1b[0m"
	padding := 6 - len(method)
	return fmt.Sprintf("[ %s%s%s ]%*s [ %s%d%s ] %s",
		MethodColor(method), method, resetColor, padding, "",
		StatusColor(status), status, resetColor,
		msg)
}

func createRequestConsoleEncoderConfig() zapcore.EncoderConfig {
	levelEncoder, timeEncoder := createCustomEncoders(true)
	return zapcore.EncoderConfig{
		TimeKey:          "timestamp",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "msg",
		EncodeLevel:      levelEncoder,
		EncodeTime:       timeEncoder,
		ConsoleSeparator: " ",
	}
}

func createRequestLogger() *zap.Logger {
	encoderConfig := createRequestConsoleEncoderConfig()
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	core := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	return zap.New(core)
}

func RequestInfoLog(msg, method string, status int) {
	formattedMsg := formatRequestMessage(msg, method, status)
	requestLogger.Info(formattedMsg)
}

func RequestDebugLog(msg, method string, status int) {
	formattedMsg := formatRequestMessage(msg, method, status)
	requestLogger.Debug(formattedMsg)
}

func RequestWarnLog(msg, method string, status int) {
	formattedMsg := formatRequestMessage(msg, method, status)
	requestLogger.Warn(formattedMsg)
}

func RequestErrorLog(msg, method string, status int) {
	formattedMsg := formatRequestMessage(msg, method, status)
	requestLogger.Error(formattedMsg)
}

func RequestFatalLog(msg, method string, status int) {
	formattedMsg := formatRequestMessage(msg, method, status)
	requestLogger.Fatal(formattedMsg)
}
