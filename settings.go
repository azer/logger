package logger

import (
	"errors"
	"os"
	"strings"
)

var (
	out                 = os.Stderr
	verbosity           = initVerbosity()
	enabled, allEnabled = initEnabled()
)

// Enabled - Returns the map representing enabled logs, and whether all loggers are
// enabled regardless of name. These are initialised from environment variables
// at start-up.
func Enabled() (map[string]bool, bool) {
	return enabled, allEnabled
}

// initEnabled - Gets the LOG environment variable, splits on commas, and stores the
// values as loggers that are permitted to print. Does not validate that loggers
// exist; typos may break loggers silently.
func initEnabled() (map[string]bool, bool) {
	val := os.Getenv("LOG")

	if val == "*" {
		return nil, true
	}

	all := map[string]bool{}
	keys := strings.Split(val, ",")

	for _, key := range keys {
		// Trim in case people pass space-gapped names.
		all[strings.TrimSpace(key)] = true
	}

	return all, false
}

// IsEnabled - Returns whether a given logger is present and enabled.
func IsEnabled(name string) bool {
	if allEnabled {
		return true
	}

	value, ok := enabled[name]
	if !ok {
		return false
	}
	return value
}

// SetEnabled - Enable or Disable loggers by name.
func SetEnabled(name string, state bool) {
	enabled[name] = state
}

// SetAllEnabled - Enable or Disable *all* loggers.
// NB: This does not change the status of individual loggers, only whether *all*
// loggers should be allowed print. In other words, setting this to 'false' does
// not disable loggers that are permitted to print, only those that are not.
func SetAllEnabled(state bool) {
	allEnabled = state
}

// Populates default value of 'verbosity' variable from Envvar.
func initVerbosity() int {
	switch strings.ToUpper(os.Getenv("LOG_LEVEL")) {
	case "ERROR":
		return 3
	case "TIMER":
		return 2
	default:
		return 1
	}
}

// Verbosity - Get the value of the LOG_LEVEL environment variable, converted to an integer.
// Default/Invalid value is 1. If "TIMER", this is 2. If "ERROR", this is 3.
func Verbosity() int {
	return verbosity
}

// SetVerbosity - Manually set the logger verbosity; must be one of:
// * 1 - Info, Timer and Error
// * 2 - Timer and Error
// * 3 - Error only
func SetVerbosity(level int) error {
	if level < 1 && level > 3 {
		return errors.New("Verbosity can only be set to a value of 1, 2 or 3")
	}
	verbosity = level
	return nil
}

// SetOutput directs logs to a file; this disables terminal colors
func SetOutput(w *os.File) {
	colorEnabled = false
	out = w
}
