package logger

import (
	"fmt"
)

// Logger is the unit of the logger package, a smart, pretty-printing gate between
// the program and the output stream.
type Logger struct {
	// Name by which the logger is identified when enabling or disabling it, and by envvar.
	Name string
	// Color is chosen automatically. It is a terminal escape code, not a color name.
	Color string
}

// New returns a logger bound to the given name. It will only be active when
// the name is enabled in the global Enabled map, which can be done using the
// LOG=<names> environment variable, or by manually setting it true in the Enabled map.
// At time of creation this adds a false value for its name to Enabled if no value
// already exists; in other words, it defaults to disabled. This may cause a
// race condition in Enabled, however, if another goroutine is accessing the Map.
func New(name string) *Logger {
	// Ensure there's an entry, but if not yet true then no env-var has yet enabled
	// it so default to disabled.
	if _, ok := Enabled[name]; !ok {
		Enabled[name] = false
	}
	return &Logger{
		Name:  name,
		Color: nextColor(),
	}
}

// IsEnabled checks whether this logger is enabled in the global roster
// (logger.Enabled map), or enabled due to logger.AllEnabled being set to true.
func (l *Logger) IsEnabled() bool {
	if AllEnabled {
		return true
	}

	value, ok := Enabled[l.Name]
	if !ok {
		return false
	}
	return value
}

// Info prints log information to the screen that is informational in nature; it
// is the most verbose level (1) and is disabled at verbosity levels 2 and 3.
func (l *Logger) Info(format string, v ...interface{}) {
	if Verbosity > 1 {
		return
	}

	if !l.IsEnabled() {
		return
	}

	v, attrs := SplitAttrs(v...)

	l.Output(1, "INFO", fmt.Sprintf(format, v...), attrs)
}

// Timer returns a timer sub-logger. Timers have verbosity level 2, they are
// considered a higher priority than Info and lower than Error.
// If a logger was disabled when a timer is created, the timer remains disabled
// even if the logger is enabled during the timer's lifespan.
func (l *Logger) Timer() *Timer {
	return &Timer{
		Logger:    l,
		Start:     Now(),
		IsEnabled: l.IsEnabled() && Verbosity < 3,
	}
}

// Error logs an error message using fmt. It has log-level 3, the highest level.
func (l *Logger) Error(format string, v ...interface{}) {
	if !l.IsEnabled() {
		return
	}

	v, attrs := SplitAttrs(v...)

	l.Output(3, "ERROR", fmt.Sprintf(format, v...), attrs)
}

// Output is the lower-level call delegated to by Info/Timer/Error, and can be used
// to directly write to the underlying buffer regardless of log-level.
func (l *Logger) Output(verbosity int, sort string, msg string, attrs *Attrs) {
	l.Write(l.Format(verbosity, sort, msg, attrs))
}

func (l *Logger) Write(log string) {
	fmt.Fprintln(out, log)
}
