package logger_test

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/wizodds/go-logger"
)

func TestLogger(t *testing.T) {
	log := logger.NewLogger(logger.LoggerOptions{})
	log.Info("This is a info log.")
}

func TestLoggerWithOptions(t *testing.T) {
	opts := logger.LoggerOptions{
		Level:      "debug",
		Name:       "DEMO",
		HideCaller: false,
	}
	log := logger.NewLogger(opts)
	log.Debug("This is a debug log.")
	log.Info("This is a info log.")
	log.Warn("This is a warn log.")
	log.Error("This is a error log.")
}

func TestToLevel(t *testing.T) {
	testCases := []struct {
		level    string
		expected zapcore.Level
	}{
		{level: "debug", expected: zap.DebugLevel},
		{level: "info", expected: zap.InfoLevel},
		{level: "warn", expected: zap.WarnLevel},
		{level: "error", expected: zap.ErrorLevel},
		{level: "dpanic", expected: zap.DPanicLevel},
		{level: "panic", expected: zap.PanicLevel},
		{level: "fatal", expected: zap.FatalLevel},
		{level: "", expected: zap.InfoLevel},
	}
	for _, tc := range testCases {
		if r := logger.ToLevel(tc.level); r != tc.expected {
			t.Fatalf("expect %v but got %v", tc.expected, r)
		}
	}
}
