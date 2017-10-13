package logger

import (
	"fmt"
	"time"
)

type Attrs map[string]interface{}

type Log struct {
	Package     string `json:"package"`
	Level       string `json:"level"`
	Message     string `json:"msg"`
	Attrs       *Attrs `json:"attrs"`
	Time        int64  `json:"time"`
	Elapsed     int64  `json:"elapsed"`
	ElapsedNano int64  `json:"elapsed_nano"`
}

func (log *Log) End(msg string, args ...interface{}) {
	v, attrs := SplitAttrs(args)
	elapsed := Now() - log.Time

	log.Attrs = attrs
	log.Elapsed = elapsed / 1000000
	log.ElapsedNano = elapsed
	log.Message = fmt.Sprintf(msg, v...)

	runtime.Log(log)
}

// SplitAttrs checks if the last item passed in v is an Attrs instance,
// if so it returns it separately. If not, v is returned as-is with a nil Attrs.
func SplitAttrs(v []interface{}) ([]interface{}, *Attrs) {
	if len(v) == 0 {
		return v, nil
	}

	attrs, ok := v[len(v)-1].(Attrs)

	if !ok {
		return v, nil
	}

	v = v[:len(v)-1]
	return v, &attrs
}

// Now is a shortcut for returning the current time in Unix nanoseconds.
func Now() int64 {
	return time.Now().UnixNano()
}
