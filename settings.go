package logger

import (
	"os"
	"strings"
)

var (
	out = os.Stderr
	verbosity = Verbosity()
	enabled, allEnabled = Enabled()
)

func Enabled () (map[string]bool, bool) {
	val := os.Getenv("LOG")

	if val == "*" {
		return nil, true
	}

	all := map[string]bool{}
	keys := strings.Split(val, ",")

	for _, key := range keys {
		all[key] = true
	}

	return all, false
}

func IsEnabled (name string) bool {
	if allEnabled {
		return true
	}

	_, ok := enabled[name]
	return ok
}

func Verbosity () int {
	level := os.Getenv("LOG_LEVEL")

	if strings.ToUpper(level) == "TIMER" {
		return 2
	}

	if strings.ToUpper(level) == "ERROR" {
		return 3
	}

	return 1
}
