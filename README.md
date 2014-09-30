## logger

Minimalistic logging library for Go.

## Install

```go
$ go get github.com/azer/logger
```

## Manual

First create an instance with a preferred name;

```go
import "github.com/azer/logger"

var log = logger.New("example-app")

func main () {
  port := 8080
  log.Info("Running at %d", port)

  err := DoSomething()

  if err != nil {
    log.Error("Failed: %s", err.Error())
  }
}
```

If it's running on a terminal, it outputs like this;

If not;
