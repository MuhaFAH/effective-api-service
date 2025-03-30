package logger

import (
	log "github.com/charmbracelet/log"
	"os"
	"strings"
)

type MinimalisticLogger struct {
	*log.Logger
}

func NewMinimalLogger(level string) MinimalisticLogger {
	logger := log.New(os.Stdout)

	logger.SetLevel(getLevel(strings.ToLower(level)))

	return MinimalisticLogger{logger}
}

func getLevel(level string) log.Level {
	switch level {
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	default:
		return log.InfoLevel
	}
}
