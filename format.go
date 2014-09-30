package logger

import (
	"github.com/azer/is-terminal"
	"fmt"
	"time"
	"syscall"
)

var colorEnabled = isterminal.IsTerminal(syscall.Stderr)

var (
	colorIndex = 0
	grey = "\x1B[90m"
	white = "\033[37m"
	reset = "\033[0m"
	bold = "\033[1m"
	red = "\033[31m"
	cyan = "\033[36m"
	colors = []string{
		"\033[34m", // blue
		"\033[32m", // green
		"\033[36m", // cyan
		"\033[33m", // yellow
		"\033[35m", // magenta
	}
)

func (l *Logger) Format (verbosity int, sort string, msg string) string {
	t := time.Now()

	if !colorEnabled {
		return fmt.Sprintf("time=\"%s\" package=\"%s\" level=\"%s\" msg=\"%s\"", t, l.Name, sort, msg)
	}

	tf := t.Format("01.02.06 15:04:05.000")

	if verbosity == 3 {
		return fmt.Sprintf("%s%s %s%s%s(%s!%s)%s:%s %s", grey, tf, l.Color, l.Name, grey, red, grey, l.Color, reset, msg)
	}

	return fmt.Sprintf("%s%s %s%s:%s %s", grey, tf, l.Color, l.Name, reset, msg)
}

func nextColor () string {
	colorIndex = colorIndex + 1
	return colors[colorIndex % len(colors)]
}
