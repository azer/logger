package logger

import (
	"fmt"
	"time"
)

type Timer struct {
	Logger    *Logger
	Start     int64
	IsEnabled bool
}

func (t *Timer) End(format string, v ...interface{}) {
	if !t.IsEnabled {
		return
	}

	elapsed := Now() - t.Start

	v, attrs := SplitAttrs(v...)
	t.Logger.Write(t.Format(elapsed, fmt.Sprintf(format, v...), attrs))
}

func (t *Timer) Format(elapsed int64, msg string, customAttrs *Attrs) string {
	if !colorEnabled {
		elapsedMS := elapsed / 1000000
		attrs := fmt.Sprintf(" \"elapsed\": %d, \"elapsed_nano\": %d,", elapsedMS, elapsed)

		if strCustomAttrs := t.Logger.JSONFormatAttrs(customAttrs); len(strCustomAttrs) > 0 {
			attrs = fmt.Sprintf("%s %s", attrs, strCustomAttrs)
		}

		return t.Logger.JSONFormat("TIMER", msg, attrs)
	}

	prefix := fmt.Sprintf("%s(%s%s%s)%s", grey, reset, time.Duration(elapsed), grey, t.Logger.Color)

	return t.Logger.ColorfulFormat(prefix, msg, customAttrs)
}

func Now() int64 {
	return time.Now().UnixNano()
}
