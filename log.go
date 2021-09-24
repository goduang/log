//
// Package log defines some simple functions for logging.
//
// In your main.go file, use `log.SetLogger(&log.Config{})` to initialize the logger.
//
// In other go files, use `log.Info("the log message", "key", "value")` to print the log messages.
//
// Example:
//
//	log.SetLogger(&log.Config{
//		Level: "debug",
//		Format: "text",
//		Layout: "2006-01-02T15:04:05.000000Z",
//		NoCaller: false,
//	})
//
//	log.Info("the info log message", "key", "value")
//	log.Debug("the debug log message")
//
//	ctrLog := log.With("name", "controller")
//	ctrLog.Info("the info log message", "key", "value")
//	ctrLog.Debug("the debug log message")
//
package log

import (
	"fmt"
	"os"
	"time"

	golog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

var logging loggingT

type loggingT struct {
	logger golog.Logger
}

// Config contains all the logger settings
type Config struct {
	Level    string
	Format   string
	Layout   string
	NoCaller bool
}

func configDefaulter(config *Config) {
	if config.Level == "" {
		config.Level = "info"
	}

	if config.Format == "" {
		config.Format = "text"
	}

	if config.Layout == "" {
		config.Layout = "2006-01-02T15:04:05.000000Z"
	}
}

// SetLogger initializes a logger with the given configuration
func SetLogger(config *Config) {
	if config == nil {
		config = &Config{}
	}
	configDefaulter(config)
	levelOption := level.AllowInfo()
	switch config.Level {
	case "all":
		levelOption = level.AllowAll()
	case "debug":
		levelOption = level.AllowDebug()
	case "info":
		levelOption = level.AllowInfo()
	case "warn":
		levelOption = level.AllowWarn()
	case "error":
		levelOption = level.AllowError()
	case "none":
		levelOption = level.AllowNone()
	default:
		fmt.Printf("unrecognized log level %q\n", config.Level)
		os.Exit(1)
	}

	var logger golog.Logger
	w := golog.NewSyncWriter(os.Stderr)
	if config.Format == "text" {
		logger = golog.NewLogfmtLogger(w)
	} else if config.Format == "json" {
		logger = golog.NewJSONLogger(w)
	} else {
		fmt.Printf("unrecognized log format %q\n", config.Format)
		os.Exit(1)
	}

	logger = golog.With(logger, "ts", golog.TimestampFormat(time.Now, config.Layout))
	logger = level.NewFilter(logger, levelOption)

	if !config.NoCaller {
		logger = golog.With(logger, "caller", golog.Caller(4))
	}

	logging.logger = logger
}

// With returns a new contextual logger with keyvals prepended
func With(keyvals ...interface{}) loggingT {
	return loggingT{logger: golog.With(logging.logger, keyvals...)}
}

// Debug prints a msg and Key/DebugValue pair
func (l *loggingT) Debug(msg string, keyvals ...interface{}) {
	keyvals = prepend(msg, keyvals...)
	level.Debug(l.logger).Log(keyvals...)
}

// Info prints a msg and Key/InfoValue pair
func (l *loggingT) Info(msg string, keyvals ...interface{}) {
	keyvals = prepend(msg, keyvals...)
	level.Info(l.logger).Log(keyvals...)
}

// Warn prints a msg and Key/WarnValue pair
func (l *loggingT) Warn(msg string, keyvals ...interface{}) {
	keyvals = prepend(msg, keyvals...)
	level.Warn(l.logger).Log(keyvals...)
}

// Error prints a msg and Key/ErrorValue pair
func (l *loggingT) Error(msg string, keyvals ...interface{}) {
	keyvals = prepend(msg, keyvals...)
	level.Error(l.logger).Log(keyvals...)
}

// Fatal prints a msg and Key/ErrorValue pair and exit
func (l *loggingT) Fatal(msg string, keyvals ...interface{}) {
	keyvals = prepend(msg, keyvals...)
	level.Error(l.logger).Log(keyvals...)
	os.Exit(1)
}

// Debug prints a msg and Key/DebugValue pair
func Debug(msg string, keyvals ...interface{}) {
	keyvals = prepend(msg, keyvals...)
	level.Debug(logging.logger).Log(keyvals...)
}

// Info prints a msg and Key/InfoValue pair
func Info(msg string, keyvals ...interface{}) {
	keyvals = prepend(msg, keyvals...)
	level.Info(logging.logger).Log(keyvals...)
}

// Warn prints a msg and Key/WarnValue pair
func Warn(msg string, keyvals ...interface{}) {
	keyvals = prepend(msg, keyvals...)
	level.Warn(logging.logger).Log(keyvals...)
}

// Error prints a msg and Key/ErrorValue pair
func Error(msg string, keyvals ...interface{}) {
	keyvals = prepend(msg, keyvals...)
	level.Error(logging.logger).Log(keyvals...)
}

// Fatal prints a msg and Key/ErrorValue pair and exit
func Fatal(msg string, keyvals ...interface{}) {
	keyvals = prepend(msg, keyvals...)
	level.Error(logging.logger).Log(keyvals...)
	os.Exit(1)
}

func prepend(msg string, keyvals ...interface{}) []interface{} {
	kvs := make([]interface{}, len(keyvals)+2)
	kvs[0], kvs[1] = "msg", msg
	copy(kvs[2:], keyvals)

	return kvs
}
