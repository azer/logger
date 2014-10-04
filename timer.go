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

	t.Logger.Write(t.Format(Now()-t.Start, fmt.Sprintf(format, v...)))
}

func (t *Timer) Format(elapsed int64, msg string) string {
	now := time.Now()

	if !colorEnabled {
		return fmt.Sprintf("{ \"time\":\"%s\", \"package\":\"%s\", \"level\":\"TIMER\", \"elapsed\":\"%d\", \"msg\":\"%s\" }", now, t.Logger.Name, elapsed, msg)
	}

	tf := now.Format("01.02.06 15:04:05.000")

	return fmt.Sprintf("%s%s %s%s%s (%s%s%s)%s:%s %s", grey, tf, t.Logger.Color, t.Logger.Name, grey, reset, time.Duration(elapsed), grey, t.Logger.Color, reset, msg)
}

func Now() int64 {
	return time.Now().UnixNano()
}
