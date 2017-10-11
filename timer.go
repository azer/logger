package logger

import (
	"fmt"
	"time"
)

// Timer is a special sub-logger that records the moment of its creation, and
// outputs time elapsed since that point when its End method is called.
type Timer struct {
	Logger *Logger
	Start  int64
}

// End prints the given message and the elapsed time since the Timer was created
// to the logger output. This does not print if the parent logger was disabled
// at the time of the Timer's creation, even if it is subsequently enabled.
func (t *Timer) End(format string, v ...interface{}) {
	if !t.Logger.OutputSettings().Timer {
		return
	}

	elapsed := Now() - t.Start
	//v, attrs := SplitAttrs(v...)
	//t.Logger.Write(t.Format(elapsed, fmt.Sprintf(format, v...), attrs))

	v, attrs := SplitAttrs(v...)
	if attrs == nil {
		attrs = &Attrs{}
	}

	(*attrs)["elapsed_nano"] = elapsed
	runtime.Log(t.Logger.Name, "TIMER", fmt.Sprintf(format, v...), attrs)
}

// Now is a shortcut for returning the current time in Unix nanoseconds.
func Now() int64 {
	return time.Now().UnixNano()
}
