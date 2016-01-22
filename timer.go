package logger

import (
	"fmt"
	"time"
)

// Timer is a special sub-logger that records the moment of its creation, and
// outputs time elapsed since that point when its End method is called.
type Timer struct {
	Logger    *Logger
	Start     int64
	IsEnabled bool
}

// End prints the given message and the elapsed time since the Timer was created
// to the logger output. This does not print if the parent logger was disabled
// at the time of the Timer's creation, even if it is subsequently enabled.
func (t *Timer) End(format string, v ...interface{}) {
	if !t.IsEnabled {
		return
	}

	elapsed := Now() - t.Start

	v, attrs := SplitAttrs(v...)
	t.Logger.Write(t.Format(elapsed, fmt.Sprintf(format, v...), attrs))
}

// Format is used to create a nicely formatted timestamp and message for Timer.
func (t *Timer) Format(elapsed int64, msg string, customAttrs *Attrs) string {
	if !colorEnabled {
		elapsedMS := elapsed / 1000000
		attrs := fmt.Sprintf(" \"elapsed\": %d, \"elapsed_nano\": %d,", elapsedMS, elapsed)

		if strCustomAttrs := t.Logger.JSONFormatAttrs(customAttrs); len(strCustomAttrs) > 0 {
			attrs = fmt.Sprintf("%s %s", attrs, strCustomAttrs)
		}

		return t.Logger.JSONFormat("TIMER", msg, attrs)
	}

	prefix := fmt.Sprintf("(%s%s)%s", reset, time.Duration(elapsed), t.Logger.Color)

	return t.Logger.PrettyFormat(prefix, msg, customAttrs)
}

// Now is a shortcut for returning the current time in Unix nanoseconds.
func Now() int64 {
	return time.Now().UnixNano()
}
