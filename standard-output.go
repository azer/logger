package logger

import (
	"fmt"
	"github.com/azer/is-terminal"
	"os"
	"strings"
	"syscall"
	"time"
)

type StandardWriter struct {
	ColorsEnabled bool
}

func (sw StandardWriter) Write(name, sort, msg string, attrs *Attrs) {
	fmt.Fprintln(os.Stderr, sw.Format(name, sort, msg, attrs))
}

func (sw *StandardWriter) Format(name, sort, msg string, attrs *Attrs) string {
	if sw.ColorsEnabled {
		return sw.PrettyFormat(name, sort, msg, attrs)
	} else {
		return sw.JSONFormat(name, sort, msg, attrs)
	}
}

func (standardWriter *StandardWriter) JSONFormat(name, sort, msg string, attrs *Attrs) string {
	return ""
}

func (sw *StandardWriter) PrettyFormat(name, sort, msg string, attrs *Attrs) string {
	return fmt.Sprintf("%s %s %s%s",
		time.Now().Format("15:04:05.000"),
		sw.PrettyLabel(name, sort, attrs),
		msg,
		sw.PrettyAttrs(attrs))
}

func (sw *StandardWriter) PrettyAttrs(attrs *Attrs) string {
	if attrs == nil {
		return ""
	}

	result := ""
	for key, val := range *attrs {
		if key != "elapsed_nano" {
			result = fmt.Sprintf("%s %s=%v", result, key, val)
		}
	}

	return result
}

func (sw *StandardWriter) PrettyLabel(name, sort string, attrs *Attrs) string {
	return fmt.Sprintf("%s%s%s:%s",
		colorFor(name),
		strings.Title(strings.ToLower(sort)),
		sw.PrettyLabelExt(name, sort, attrs),
		reset)
}

func (sw *StandardWriter) PrettyLabelExt(name, sort string, attrs *Attrs) string {
	if sort == "ERROR" {
		return fmt.Sprintf("(%s!%s)", red, colorFor(name))
	}

	if sort == "TIMER" {
		elapsed := "?"
		if elapsedRaw, ok := (*attrs)["elapsed_nano"]; ok {
			if elapsedNano, ok := elapsedRaw.(int64); ok {
				elapsed = fmt.Sprintf("%v", time.Duration(elapsedNano))
			}
		}

		return fmt.Sprintf("(%s%s%s)", reset, elapsed, colorFor(name))
	}

	return ""
}

func NewStandardOutput() (OutputWriter, map[string]*OutputSettings) {
	var writer OutputWriter = StandardWriter{
		ColorsEnabled: isterminal.IsTerminal(syscall.Stderr),
	}

	defaultOutputSettings := parseVerbosityLevel(os.Getenv("LOG_LEVEL"))
	return writer, parsePackageSettings(os.Getenv("LOG"), defaultOutputSettings)
}

// Accepts: foo,bar,qux@timer
//          *
//          *@error
//          *@error,database@timer
func parsePackageSettings(input string, defaultOutputSettings *OutputSettings) map[string]*OutputSettings {
	all := map[string]*OutputSettings{}
	items := strings.Split(input, ",")

	for _, item := range items {
		name, verbosity := parsePackageName(item)
		if verbosity == nil {
			verbosity = defaultOutputSettings
		}

		all[name] = verbosity
	}

	return all
}

// Accepts: users
//          database@timer
//          server@error
func parsePackageName(input string) (string, *OutputSettings) {
	parsed := strings.Split(input, "@")
	name := strings.TrimSpace(parsed[0])

	if len(parsed) > 1 {
		return name, parseVerbosityLevel(parsed[1])
	}

	return name, nil
}

func parseVerbosityLevel(val string) *OutputSettings {
	val = strings.ToUpper(strings.TrimSpace(val))

	if val == "MUTE" {
		return &OutputSettings{}
	}

	s := &OutputSettings{
		Info:  true,
		Timer: true,
		Error: true,
	}

	if val == "TIMER" {
		s.Info = false
	}

	if val == "ERROR" {
		s.Info = false
		s.Timer = false
	}

	return s
}
