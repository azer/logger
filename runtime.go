package logger

import (
	"os"
)

var (
	runtime *Runtime
	muted   = &OutputSettings{}
	verbose = &OutputSettings{
		Info:  true,
		Timer: true,
		Error: true,
	}
)

func init() {
	runtime = &Runtime{
		Writers: []OutputWriter{
			NewStandardOutput(os.Stderr),
		},
	}
}

type OutputWriter interface {
	Write(log *Log)
}

type OutputSettings struct {
	Info  bool
	Timer bool
	Error bool
}

type Runtime struct {
	Writers []OutputWriter
}

func (runtime *Runtime) Log(log *Log) {
	if len(runtime.Writers) == 0 {
		return
	}

	// Avoid getting into a loop if there is just one writer
	if len(runtime.Writers) == 1 {
		runtime.Writers[0].Write(log)
	} else {
		for _, w := range runtime.Writers {
			w.Write(log)
		}
	}
}

// Add a new writer
func Hook(writer OutputWriter) {
	runtime.Writers = append(runtime.Writers, writer)
}

// Legacy method
func SetOutput(file *os.File) {
	writer := NewStandardOutput(file)
	runtime.Writers[0] = writer
}
