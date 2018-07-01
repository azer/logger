package main

import (
	"fmt"
	"github.com/azer/logger"
	"time"
)

var log = logger.New("app")

type CustomWriter struct{}

func (customWriter *CustomWriter) Init() {

}

func (cw CustomWriter) Write(log *logger.Log) {
	fmt.Println("custom log -> ", log.Package, log.Level, log.Message, log.Attrs)
}

func main() {
	logger.Hook(&CustomWriter{})
	log.Info("he-yo")

	log.Info("Requesting an image", logger.Attrs{
		"file": "foo/bar.jpg",
	})

	timer := log.Timer()
	time.Sleep(time.Millisecond * 250)
	timer.End("Fetched foo/bar.jpg")

	log.Error("Failed, shutting down...")
}
