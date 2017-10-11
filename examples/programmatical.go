package main

import (
	"fmt"
	"github.com/azer/logger"
	"time"
)

var log = logger.New("app")

type CustomWriter struct{}

func (cw CustomWriter) Write(pkg, sort, msg string, attrs *logger.Attrs) {
	fmt.Println("custom log -> ", pkg, sort, msg, attrs)
}

func main() {
	logger.Hook(&CustomWriter{})
	log.Info("he-yo")

	log.Info("Requesting an image at foo/bar.jpg")
	timer := log.Timer()
	time.Sleep(time.Millisecond * 250)
	timer.End("Fetched foo/bar.jpg")

	log.Error("Failed to start, shutting down...")
}
