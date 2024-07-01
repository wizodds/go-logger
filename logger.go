package logger

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerOptions struct {
	Level      string
	Name       string
	HideCaller bool
}

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(opts LoggerOptions) *Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true
	cfg.DisableCaller = opts.HideCaller
	cfg.EncoderConfig.ConsoleSeparator = " "
	cfg.EncoderConfig.EncodeTime = localTimeEncoder
	cfg.EncoderConfig.EncodeCaller = localCallerEncoder
	cfg.EncoderConfig.EncodeLevel = localLevelEncoder(opts)
	cfg.Level = zap.NewAtomicLevelAt(ToLevel(opts.Level))
	logger, _ := cfg.Build()
	defer logger.Sync()
	return &Logger{SugaredLogger: logger.Sugar()}
}

func localTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000 -07:00"))
}

func localCallerEncoder(c zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	f := strings.Split(c.Function, "/")
	enc.AppendString(fmt.Sprintf("[%s %s]", c.TrimmedPath(), f[len(f)-1]))
}

func localLevelEncoder(opts LoggerOptions) zapcore.LevelEncoder {
	emojiLogLevels := map[string]string{
		"debug":  "🟪 DEBUG ",
		"info":   "⬜️ INFO  ",
		"warn":   "🟧 WARN  ",
		"error":  "🟥 ERROR ",
		"dpanic": "🟥 DPANIC",
		"panic":  "🟥 PANIC ",
		"fatal":  "🟥 FATAL ",
	}
	name := ""
	if len(opts.Name) > 0 {
		name = fmt.Sprintf(" [%s]", opts.Name)
	}
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("[go] %s%s", emojiLogLevels[l.String()], name))
	}
}

func ToLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
