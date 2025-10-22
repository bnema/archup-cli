package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

var (
	// Logger is the global logger instance
	Logger *log.Logger
)

// Init initializes the global logger with appropriate settings
func Init(debug bool) {
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    debug,
		ReportTimestamp: true,
		TimeFormat:      "15:04:05",
	})

	if debug {
		Logger.SetLevel(log.DebugLevel)
	} else {
		Logger.SetLevel(log.InfoLevel)
	}
}

// Info logs an info message
func Info(msg string, keyvals ...interface{}) {
	if Logger != nil {
		Logger.Info(msg, keyvals...)
	}
}

// Debug logs a debug message
func Debug(msg string, keyvals ...interface{}) {
	if Logger != nil {
		Logger.Debug(msg, keyvals...)
	}
}

// Warn logs a warning message
func Warn(msg string, keyvals ...interface{}) {
	if Logger != nil {
		Logger.Warn(msg, keyvals...)
	}
}

// Error logs an error message
func Error(msg string, keyvals ...interface{}) {
	if Logger != nil {
		Logger.Error(msg, keyvals...)
	}
}

// Fatal logs a fatal message and exits
func Fatal(msg string, keyvals ...interface{}) {
	if Logger != nil {
		Logger.Fatal(msg, keyvals...)
	}
}
