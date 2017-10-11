## logger

Minimalistic logging library for Go. Supports timers, filtering by package names, log levels and structured output.

![](https://i.cloudup.com/rUyno2tHCx.png)

## Install

```bash
$ go get github.com/azer/logger
```

## Manual

First create an instance with a preferred name;

```go
import "github.com/azer/logger"

var log = logger.New("example-app")
```

It gives only three methods; `Info`, `Timer` and `Error`.

```go
log.Info("Running at %d", 8080)

err := DoSomething()

if err != nil {
  log.Error("Failed: %s", err.Error())
}
```

Check out [example/simple.go](https://github.com/azer/logger/blob/master/example/simple.go) for a more detailed example.

### Filtering

By default, it won't output. Enable it with `LOG` environment variable:

```bash
$ LOG=* go run example-app.go
```

This will enable all logs. You can choose the packages that you'd like to display:

```bash
$ LOG=images,users go run example-app.go
```

In the above example, you'll only see logs from `images` and `users` packages. You can allow everything except some specific packages using `LOG_EXCEPT` parameter:

```bash
$ LOG=* LOG_EXCEPT=users go run example-app.go
```


You can filter logs by their level, too. If `INFO` level is not useful for your case, pass `LOG_LEVEL`:

```bash
$ LOG=images,users LOG_LEVEL=error go run example-app.go
```

### Timers

You can use timer logs for measuring your program. For example;

```go
timer := log.Timer()

image, err := PullImage("http://foo.com/bar.jpg")

timer.End("Fetched foo.com/bar.jpg")
```

Timer log lines will be outputting the elapsed time in time.Duration in a normal terminal, or in int64 format when your program is running on a non-terminal environment.
See below documentation for more info.

### Structured Output

When your app isn't running on a terminal, it'll change the output in JSON:

```
{ "time":"2014-10-04 11:44:22.418595705 -0700 PDT", "package":"database", "level":"INFO", "msg":"Connecting to mysql://azer@localhost:9900/foobar" }
{ "time":"2014-10-04 11:44:22.418600851 -0700 PDT", "package":"images", "level":"INFO", "msg":"Requesting an image at foo/bar.jpg" }
{ "time":"2014-10-04 11:44:22.668645527 -0700 PDT", "package":"images", "level":"TIMER", "elapsed":"250032416", "msg":"Fetched foo/bar.jpg" }
{ "time":"2014-10-04 11:44:22.668665527 -0700 PDT", "package":"database", "level":"ERROR", "msg":"Fatal connection error." }
{ "time":"2014-10-04 11:44:22.668673037 -0700 PDT", "package":"users", "level":"INFO", "msg":"John just logged  from Istanbul" }
{ "time":"2014-10-04 11:44:22.668676732 -0700 PDT", "package":"websocket", "level":"INFO", "msg":"Connecting..." }
{ "time":"2014-10-04 11:44:22.668681092 -0700 PDT", "package":"websocket", "level":"ERROR", "msg":"Unable to connect." }
{ "time":"2014-10-04 11:44:22.919726985 -0700 PDT", "package":"app", "level":"ERROR", "msg":"Failed to start, shutting down..." }
```

So you can parse & process the output easily. Here is a command that lets you see the JSON output in your terminal;

```
LOG=* go run example/simple.go 2>&1 | less
```

### Attributes

To add custom attributes to the structured output;

```go
log.Info("Sending an e-mail...", logger.Attrs{
  "from": "foo@bar.com",
  "to": "qux@corge.com",
})
```

The above log will appear in the structured output as:

```go
{ "time":"2014-10-04 11:44:22.919726985 -0700 PDT", "package":"mail", "level":"INFO", "msg":"Sending an e-mail", "from": "foo@foobar.com", "to": "qux@corge.com" }
```

In your command-line as:

![](https://cldup.com/n4Uia8v1uo.png)

### Setting The Output

By default, it outputs to `stderr`. You can change it by calling `SetOutput` with an `*os.File` parameter.
