package logger

import (
	"fmt"
	"github.com/azer/is-terminal"
	"syscall"
	"time"
)

var colorEnabled = isterminal.IsTerminal(syscall.Stderr)

var (
	colorIndex = 0
	grey       = "\x1B[90m"
	white      = "\033[37m"
	reset      = "\033[0m"
	bold       = "\033[1m"
	red        = "\033[31m"
	cyan       = "\033[36m"
	colors     = []string{
		"\033[34m", // blue
		"\033[32m", // green
		"\033[36m", // cyan
		"\033[33m", // yellow
		"\033[35m", // magenta
	}
)

func (l *Logger) Format(verbosity int, sort string, msg string, attrs *Attrs) string {
	if !colorEnabled {
		return l.JSONFormat(sort, msg, l.JSONFormatAttrs(attrs))
	}

	return l.ColorfulFormat(l.ColorfulPrefix(verbosity), msg, attrs)
}

func (l *Logger) JSONFormat(sort string, msg string, attrs string) string {
	return fmt.Sprintf("{ \"time\":\"%s\", \"package\":\"%s\", \"level\":\"%s\",%s \"msg\":\"%s\" }", time.Now(), l.Name, sort, attrs, msg)
}

func (l *Logger) JSONFormatAttrs (attrs *Attrs) string {
	result := ""

	if attrs == nil {
		return ""
	}

	for key, val := range *attrs {
		if val, ok := val.(int); ok {
			result = fmt.Sprintf("%s \"%s\": %d,", result, key, val)
			continue
		}

		result = fmt.Sprintf("%s \"%s\":\"%s\",", result, key, val)
	}

	return result
}

func (l *Logger) ColorfulFormat(prefix, msg string, attrs *Attrs) string {
	return fmt.Sprintf("%s%s %s%s%s:%s %s%s", grey, time.Now().Format("01.02.06 15:04:05.000"), l.Color, l.Name, prefix, reset, msg, l.ColorfulAttrs(attrs))
}

func (l *Logger) ColorfulAttrs(attrs *Attrs) string {
	result := ""
	empty := true

	if attrs == nil {
		return ""
	}

	for key, val := range *attrs {
		if empty == true {
			empty = false
		}

		if val, ok := val.(int); ok {
			result = fmt.Sprintf("%s %s=%d", result, key, val)
			continue
		}

		result = fmt.Sprintf("%s %s=%s", result, key, val)
	}

	if empty == true {
		return ""
	}

	return fmt.Sprintf("%s %s %s", grey, result, reset)
}

func (l *Logger) ColorfulPrefix(verbosity int) string {
	if verbosity != 3 {
		return ""
	}

	prefix := ""

	if verbosity == 3 {
		prefix = fmt.Sprintf("%s!", red)
	}

	return fmt.Sprintf("%s(%s%s)%s", grey, prefix, grey, l.Color)
}

func nextColor() string {
	colorIndex = colorIndex + 1
	return colors[colorIndex%len(colors)]
}
