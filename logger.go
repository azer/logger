package logger

import (
	"fmt"
)

// New returns a logger bound to the given name.
func New(name string) *Logger {
	return &Logger{
		Name: name,
	}
}

// Logger is the unit of the logger package, a smart, pretty-printing gate between
// the program and the output stream.
type Logger struct {
	// Name by which the logger is identified when enabling or disabling it, and by envvar.
	Name string
}

func (logger *Logger) Log(level, message string, args []interface{}) {
	v, attrs := SplitAttrs(args)

	runtime.Log(&Log{
		Package: logger.Name,
		Level:   level,
		Message: fmt.Sprintf(message, v...),
		Time:    Now(),
		Attrs:   attrs,
	})
}

// Info prints log information to the screen that is informational in nature.
func (l *Logger) Info(msg string, v ...interface{}) {
	l.Log("INFO", msg, v)
}

// Error logs an error message.
func (l *Logger) Error(msg string, v ...interface{}) {
	l.Log("ERROR", msg, v)
}

// Timer returns a timer sub-logger.
func (l *Logger) Timer() *Log {
	return &Log{
		Package: l.Name,
		Level:   "TIMER",
		Time:    Now(),
	}
}
