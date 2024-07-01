# go-logger

Logger for Golang

## Install

```shell
go get github.com/wizodds/go-logger
```

## Example

```go
package main

import (
  "github.com/wizodds/go-logger"
)

func main() {
  opts := logger.LoggerOptions{
    Level:      "debug",
    Name:       "DEMO",
    HideCaller: false,
  }
  log := logger.NewLogger(opts)
  log.Debug("This is a debug log.")
  log.Info("This is a info log.")
  log.Warn("This is a warn log.")
  log.Error("This is a error log.")
}
```

Output

```shell
2024-07-02 01:39:01.187 +07:00 [go] 🟪 DEBUG  [DEMO] [example/main.go:14 main.main] This is a debug log.
2024-07-02 01:39:01.188 +07:00 [go] ⬜️ INFO   [DEMO] [example/main.go:15 main.main] This is a info log.
2024-07-02 01:39:01.188 +07:00 [go] 🟧 WARN   [DEMO] [example/main.go:16 main.main] This is a warn log.
2024-07-02 01:39:01.188 +07:00 [go] 🟥 ERROR  [DEMO] [example/main.go:17 main.main] This is a error log.
```

## Reference

- [uber-go/zap][1]

## License

Licensed under the MIT License - see the [LICENSE][2] file for details.

[1]: https://github.com/uber-go/zap
[2]: https://github.com/i14t/go-logger/blob/main/LICENSE
