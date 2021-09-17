# log

A simple structured log library.

```
import "github.com/goduang/log"

log.SetLogger(nil)
log.SetLogger(&log.Config{
    Level: "debug",  // debug, info, warn, error
    Format: "text",  // text, json
    NoCaller: false, // print caller info
})

log.Debug("key", "val", "key1", "val1")
log.Info("key", "val", "key1", "val1")
log.Warn("key", "val", "key1", "val1")
log.Error("key", "val", "key1", "val1")
log.Fatal("key", "val", "key1", "val1")

log.With("name", "controller").Info("key", "value")
```
