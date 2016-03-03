package logger

import (
	"io"
	"os"
	"strings"
	"syscall"

	"github.com/azer/is-terminal"
)

var (
	// Verbosity is the level of verbosity, between 1 and 3. It is initialised from
	// an environment variable if given, you can also set this using the LevelX constants.
	Verbosity logLevelT
	// Enabled is a map of logger names to booleans indicating whether they are enabled or not.
	// It is pre-populated from environment variable "LOG", and may initially omit logs
	// not given in that envvar.
	Enabled map[string]bool
	// AllEnabled is a global boolean that activates all loggers regardless of the contents
	// of Enabled.
	AllEnabled bool

	out          io.Writer
	colorEnabled bool
)

func init() {
	out = os.Stderr
	colorEnabled = isterminal.IsTerminal(syscall.Stderr)
	Enabled, AllEnabled = initEnabled()
	Verbosity = initVerbosity()
}

type logLevelT int

const (
	// Level1 enables Info, Timer and Error log levels.
	Level1 logLevelT = iota + 1
	// Level2 enables Timer and Error log levels.
	Level2
	// Level3 enables Error log level.
	Level3
)

// initEnabled - Gets the LOG environment variable, splits on commas, and stores the
// values as loggers that are permitted to print. Does not validate that loggers
// exist; typos may break loggers silently.
func initEnabled() (map[string]bool, bool) {
	val := os.Getenv("LOG")

	if val == "*" {
		return map[string]bool{}, true
	}

	all := map[string]bool{}
	keys := strings.Split(val, ",")

	for _, key := range keys {
		// Trim in case people pass space-gapped names.
		all[strings.TrimSpace(key)] = true
	}

	return all, false
}

// Populates default value of 'verbosity' variable from Envvar.
func initVerbosity() logLevelT {
	switch strings.ToUpper(os.Getenv("LOG_LEVEL")) {
	case "ERROR":
		return Level3
	case "TIMER":
		return Level2
	default:
		return Level1
	}
}

// SetOutput directs logs to a writer other than Stderr; this disables pretty pretty-printing
// and outputs JSON instead.
func SetOutput(w io.Writer) {
	colorEnabled = false
	out = w
}
