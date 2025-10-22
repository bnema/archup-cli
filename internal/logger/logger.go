package logger

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

const (
	// LogFile is the default log file path
	LogFile = "/var/log/archup-wizard.log"
)

var (
	// Logger is the global logger instance
	Logger *log.Logger
	// logFile is the file handle for the log file
	logFile *os.File
)

// Init initializes the global logger with file and stderr output
func Init(debug bool) error {
	var writers []io.Writer

	// Always log to stderr for user feedback
	writers = append(writers, os.Stderr)

	// Try to open log file (will fail if not root, that's ok)
	f, err := os.OpenFile(LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err == nil {
		logFile = f
		writers = append(writers, f)
	}

	// Create multi-writer (stderr + file if available)
	writer := io.MultiWriter(writers...)

	Logger = log.NewWithOptions(writer, log.Options{
		ReportCaller:    debug,
		ReportTimestamp: true,
		TimeFormat:      "2006-01-02 15:04:05",
	})

	if debug {
		Logger.SetLevel(log.DebugLevel)
	} else {
		Logger.SetLevel(log.InfoLevel)
	}

	return nil
}

// Close closes the log file if it was opened
func Close() error {
	if logFile != nil {
		return logFile.Close()
	}
	return nil
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
