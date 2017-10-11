package logger

import (
	"fmt"
)

// Logger is the unit of the logger package, a smart, pretty-printing gate between
// the program and the output stream.
type Logger struct {
	// Name by which the logger is identified when enabling or disabling it, and by envvar.
	Name string
}

// New returns a logger bound to the given name.
func New(name string) *Logger {
	return &Logger{
		Name: name,
	}
}

// Returns output settings for this package. If there isn't anything specified,
// returns default muted settings.
func (logger *Logger) OutputSettings() *OutputSettings {
	if settings, ok := runtime.Settings[logger.Name]; ok {
		return settings
	}

	// If there is a "*" (Select all) setting, return that
	if settings, ok := runtime.Settings["*"]; ok {
		return settings
	}

	return muted
}

// Info prints log information to the screen that is informational in nature.
func (l *Logger) Info(format string, v ...interface{}) {
	if l.OutputSettings().Info {
		v, attrs := SplitAttrs(v...)
		runtime.Log(l.Name, "INFO", fmt.Sprintf(format, v...), attrs)
	}
}

// Timer returns a timer sub-logger.
func (l *Logger) Timer() *Timer {
	return &Timer{
		Logger: l,
		Start:  Now(),
	}
}

// Error logs an error message using fmt. It has log-level 3, the highest level.
func (l *Logger) Error(format string, v ...interface{}) {
	if l.OutputSettings().Error {
		v, attrs := SplitAttrs(v...)
		runtime.Log(l.Name, "ERROR", fmt.Sprintf(format, v...), attrs)
	}
}
