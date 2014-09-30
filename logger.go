package logger

import (
	"fmt"
)

type Logger struct {
	Name string
	IsEnabled bool
	Color string
}

func New (name string) *Logger {
	return &Logger{
		Name: name,
		IsEnabled: IsEnabled(name),
		Color: nextColor(),
	}
}

func (l *Logger) Info (format string, v ...interface{}) {
	if verbosity > 1 {
		return
	}

	if !l.IsEnabled {
		return
	}

	l.Output(1,"INFO", fmt.Sprintf(format, v...))
}

func (l *Logger) Error (format string, v ...interface{}) {
	if !l.IsEnabled {
		return
	}

	l.Output(2, "ERROR", fmt.Sprintf(format, v...))
}

func (l *Logger) Output (verbosity int, sort string, msg string) {
	fmt.Fprintln(out, l.Format(verbosity, sort, msg))
}
