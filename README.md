## logger

Minimalistic logging library for Go.

![](https://i.cloudup.com/qdQd32SXD9.png)

## Install

```bash
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

By default, it won't output. Enable it with `LOG` environment variable:

```bash
$ LOG=* go run example-app.go
```

This will enable all logs. You can choose the packages that you'd like to display:

```bash
$ LOG=images,users go run example-app.go
```

If `INFO` level is not useful for your case, pass `LOG_LEVEL`:

```bash
$ LOG=* LOG_LEVEL=error go run example-app.go
```

When your app isn't running on a terminal, it'll change the output format into:

```
time="2014-09-29 23:51:59.690196205 -0700 PDT" package="app" level="INFO" msg="Starting at 9088"
time="2014-09-29 23:51:59.690302069 -0700 PDT" package="database" level="INFO" msg="Connecting to mysql://azer@localhost:9900/foobar"
time="2014-09-29 23:51:59.690315471 -0700 PDT" package="images" level="INFO" msg="Requesting an image at /foo/bar.jpg"
time="2014-09-29 23:51:59.940415043 -0700 PDT" package="database" level="ERROR" msg="Fatal connection error."
time="2014-09-29 23:51:59.940454957 -0700 PDT" package="users" level="INFO" msg="John just logged  from Istanbul"
time="2014-09-29 23:51:59.94046777 -0700 PDT" package="websocket" level="INFO" msg="Connecting..."
time="2014-09-29 23:51:59.940476972 -0700 PDT" package="websocket" level="ERROR" msg="Unable to connect."
time="2014-09-29 23:52:00.191250959 -0700 PDT" package="app" level="ERROR" msg="Failed to start, shutting down..."
```

So you can parse & process the output easily.
