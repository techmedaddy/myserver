package utils

import (
	"log"
	"os"
)

// Logger is a simple wrapper around the standard log package.
// For small services, this is enough.
// For larger systems, you’d swap this with a structured logger like zerolog or zap.
type Logger struct {
	info  *log.Logger
	error *log.Logger
}

// NewLogger initializes a new logger instance.
func NewLogger() *Logger {
	return &Logger{
		info:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		error: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs informational messages.
func (l *Logger) Info(msg string) {
	l.info.Println(msg)
}

// Error logs error messages.
func (l *Logger) Error(err error) {
	if err != nil {
		l.error.Println(err)
	}
}
