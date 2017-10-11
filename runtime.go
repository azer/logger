package logger

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
	writer, settings := NewStandardOutput()
	runtime = &Runtime{
		Settings: settings,
		Writers:  []OutputWriter{writer},
	}
}

type OutputWriter interface {
	Write(name, sort, msg string, attrs *Attrs)
}

type OutputSettings struct {
	Info  bool
	Timer bool
	Error bool
}

type Runtime struct {
	// The verbosity level preference.
	Settings map[string]*OutputSettings
	Writers  []OutputWriter
}

func (runtime *Runtime) Log(name, sort, msg string, attrs *Attrs) {
	if len(runtime.Writers) == 0 {
		return
	}

	// Avoid getting into a loop if there is just one writer
	if len(runtime.Writers) == 1 {
		runtime.Writers[0].Write(name, sort, msg, attrs)
	} else {
		for _, w := range runtime.Writers {
			w.Write(name, sort, msg, attrs)
		}
	}
}
