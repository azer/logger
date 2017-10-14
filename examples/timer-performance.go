package main

import (
	"github.com/azer/logger"
)

func main() {

	var perf = logger.New("performance")
	var test = logger.New("test")

	timer := perf.Timer()
	for i := 0; i < 999999; i++ {
		t := test.Timer()
		t.End("foobar", logger.Attrs{
			"foo": 123,
			"bar": true,
		})
	}
	timer.End("End")
}
