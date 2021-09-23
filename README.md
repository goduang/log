# log

A simple structured log library.

```
import "github.com/goduang/log"

log.SetLogger(nil)
log.SetLogger(&log.Config{
    Level:    "debug",                      // debug, info, warn, error
    Format:   "text",                       // text, json
    NoCaller: false,                        // print caller info
    Layout:   "2006-01-02T15:04:05.000000", // the timestamp format
})

log.Debug("this is debug message", "key", "val")
log.Info("this is info message", "key", "val")
log.Warn("this is warn message", "key", "val")
log.Error("this is error message", "key", "val")
log.Fatal("this is error message", "key", "val")

log.With("name", "controller").Info("this is info message", "key", "value")
```
