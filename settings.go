package logger

/*import (
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
	// not given in that envvar. User can choose specific level for a specific package.
	Enabled map[string]logLevelT

	// AllEnabled is a global boolean that activates all loggers regardless of the contents
	// of Enabled.
	AllEnabled bool

	// In case AllEnabled is enabled and you want to filter some packages, you can pass LOG_EXCEPT=foo,bar
	Except map[string]bool

	out          io.Writer
	colorEnabled bool
)

func init() {
	Enabled, AllEnabled = initEnabled()

	if AllEnabled {
		Except = initExcept()
	}

	Verbosity = initVerbosity()
}

type logLevelT int

const (
	// Level0 mutes all logs
	Level0 logLevelT = 0
	// Level1 enables Info, Timer and Error log levels.
	Level1 logLevelT = 1
	// Level2 enables Timer and Error log levels.
	Level2 logLevelT = 2
	// Level3 enables Error log level.
	Level3 logLevelT = 3
)

// initEnabled - Gets the LOG environment variable, splits on commas, and stores the
// values as loggers that are permitted to print. Does not validate that loggers
// exist; typos may break loggers silently.
func initEnabled() (map[string]bool, bool) {
	val := os.Getenv("LOG")

	if val == "*" {
		return map[string]bool{}, true
	}

	return readPackageNames(val), false
}

func initExcept() map[string]bool {
	return readPackageNames(os.Getenv("LOG_EXCEPT"))
}

// Populates default value of 'verbosity' variable from Envvar.
func initVerbosity() logLevelT {
	return parseLogLevel(os.Getenv("LOG_LEVEL"))
}

// SetOutput directs logs to a writer other than Stderr; this disables pretty pretty-printing
// and outputs JSON instead.
func SetOutput(w io.Writer) {
	colorEnabled = false
	out = w
}
*/
