package main

import (
	"github.com/wizodds/go-logger"
)

func main() {
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
