package logger

import (
	"fmt"
	"time"
)

var (
	colorIndex = 0
	white      = "\033[37m"
	reset      = "\033[0m"
	bold       = "\033[1m"
	red        = "\033[31m"
	cyan       = "\033[36m"
	colors     = []string{
		"\033[34m", // blue
		"\033[32m", // green
		"\033[36m", // cyan
		"\033[33m", // yellow
		"\033[35m", // magenta
	}
)

// Format returns either a JSON string or a pretty line for printing to terminal,
// depending on whether logger believes it's printing to stderr or to another writer.
func (l *Logger) Format(verbosity int, sort string, msg string, attrs *Attrs) string {
	if !colorEnabled {
		return l.JSONFormat(sort, msg, l.JSONFormatAttrs(attrs))
	}

	return l.PrettyFormat(l.PrettyPrefix(verbosity), msg, attrs)
}

// JSONFormat returns a JSON string representing the data provided and some internal state
// from the logger.
func (l *Logger) JSONFormat(sort string, msg string, attrs string) string {
	// This is vulnerable to injection of invalid JSON characters.
	// For JSON formatting, json.Marshal on an anonymous struct might work?
	// Could then also dispense with JSONFormatAttrs, provided it iterates map values correctly.
	// eg:
	// jstruct := struct{
	//  Time time.Time `json:"time"`
	//  Package string `json:"package"`
	//  Level string `json:"level"`
	//  Msg map[string]interface{} `json:"msg"`
	// } {time.Now(), l.Name, sort, attrs, msg}
	// j, err := json.Marshal(jstruct)
	// if err != nil { return err }
	// return string(j)
	return fmt.Sprintf("{ \"time\":\"%s\", \"package\":\"%s\", \"level\":\"%s\",%s \"msg\":\"%s\" }", time.Now(), l.Name, sort, attrs, msg)
}

// JSONFormatAttrs converts an Attrs object to JSON.
func (l *Logger) JSONFormatAttrs(attrs *Attrs) string {
	result := ""

	if attrs == nil {
		return ""
	}

	// Vulnerable to injection of invalid characters as it doesn't perform escaping.
	// Should probably be replaced with a valid json.Marshal call?
	for key, val := range *attrs {
		if val, ok := val.(int); ok {
			result = fmt.Sprintf("%s \"%s\": %d,", result, key, val)
			continue
		}

		result = fmt.Sprintf("%s \"%s\":\"%s\",", result, key, val)
	}

	return result
}

// PrettyFormat constructs a timestamped, named, levelled log line for a given message/attrs.
func (l *Logger) PrettyFormat(prefix, msg string, attrs *Attrs) string {
	return fmt.Sprintf("%s %s%s%s:%s %s%s", time.Now().Format("15:04:05.000"), l.Color, l.Name, prefix, reset, msg, l.PrettyAttrs(attrs))
}

// PrettyAttrs formats structured data provided as Attrs for printing to terminal.
func (l *Logger) PrettyAttrs(attrs *Attrs) string {
	result := ""
	empty := true

	if attrs == nil {
		return ""
	}

	for key, val := range *attrs {
		if empty == true {
			empty = false
		}

		if val, ok := val.(int); ok {
			result = fmt.Sprintf("%s %s=%d", result, key, val)
			continue
		}

		result = fmt.Sprintf("%s %s=%s", result, key, val)
	}

	if empty == true {
		return ""
	}

	return fmt.Sprintf("%s %s", result, reset)
}

// PrettyPrefix provides a red "(!)" for Error logs only.
func (l *Logger) PrettyPrefix(verbosity int) string {
	// was one of the below (X vs. verbosity) gates supposed to be referring to
	// global verbosity, or local argument verbosity?
	if verbosity != 3 {
		return ""
	}
	return fmt.Sprintf("(%s%s)", red+"!", l.Color)
}

func nextColor() string {
	colorIndex = colorIndex + 1
	return colors[colorIndex%len(colors)]
}
