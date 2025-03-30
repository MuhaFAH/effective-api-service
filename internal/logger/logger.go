package logger

import (
	log "github.com/charmbracelet/log"
	"os"
	"strings"
)

type MinimalisticLogger struct {
	*log.Logger
}

func NewMinimalLogger(name, level string) MinimalisticLogger {
	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      "2006-01-02 15:04:05",
		Prefix:          name,
	})

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
