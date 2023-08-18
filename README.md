# log

DEPRECATED, please use standard library slog https://pkg.go.dev/log/slog instead.

A simple structured log library.

```
import "github.com/goduang/log"

log.SetLogger(nil)

log.Debug("this is debug message", "key", "val")
log.Info("this is info message", "key", "val")
log.Warn("this is warn message", "key", "val")
log.Error("this is error message", "key", "val")
log.Fatal("this is error message", "key", "val")

log.With("name", "controller").Info("this is info message", "key", "value")
```
