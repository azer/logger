package logger

import (
	"fmt"
)

type Logger struct {
	Name      string
	IsEnabled bool
	Color     string
}

func New(name string) *Logger {
	return &Logger{
		Name:      name,
		IsEnabled: IsEnabled(name),
		Color:     nextColor(),
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	if verbosity > 1 {
		return
	}

	if !l.IsEnabled {
		return
	}

	v, attrs := SplitAttrs(v...)

	l.Output(1, "INFO", fmt.Sprintf(format, v...), attrs)
}

func (l *Logger) Timer() *Timer {
	return &Timer{
		Logger:    l,
		Start:     Now(),
		IsEnabled: l.IsEnabled && verbosity < 3,
	}
}

func (l *Logger) Error(format string, v ...interface{}) {
	if !l.IsEnabled {
		return
	}

	v, attrs := SplitAttrs(v...)

	l.Output(3, "ERROR", fmt.Sprintf(format, v...), attrs)
}

func (l *Logger) Output(verbosity int, sort string, msg string, attrs *Attrs) {
	l.Write(l.Format(verbosity, sort, msg, attrs))
}

func (l *Logger) Write(log string) {
	fmt.Fprintln(out, log)
}
