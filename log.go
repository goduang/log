//
// Package log defines some simple functions for logging.
//
// In the main.go file, use `log.SetLogger(&log.Config{})` to initialize the logger.
//
// In other go files, use `log.Info("msg", "the log message")` to print the log messages.
//

package log

import (
	"fmt"
	"os"

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
	NoCaller bool
}

func configDefaulter(config *Config) {
	if config.Level == "" {
		config.Level = "info"
	}

	if config.Format == "" {
		config.Format = "text"
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

	logger = level.NewFilter(logger, levelOption)
	logger = golog.With(logger, "ts", golog.DefaultTimestamp)

	if !config.NoCaller {
		logger = golog.With(logger, "caller", golog.Caller(4))
	}

	logging.logger = logger
}

// With returns a new contextual logger with keyvals prepended
func With(keyvals ...interface{}) golog.Logger {
	return golog.With(logging.logger, keyvals...)
}

// Debug returns a logger that includes a Key/DebugValue pair
func Debug(keyvals ...interface{}) {
	level.Debug(logging.logger).Log(keyvals...)
}

// Info returns a logger that includes a Key/InfoValue pair
func Info(keyvals ...interface{}) {
	level.Info(logging.logger).Log(keyvals...)
}

// Warn returns a logger that includes a Key/WarnValue pair
func Warn(keyvals ...interface{}) {
	level.Warn(logging.logger).Log(keyvals...)
}

// Error returns a logger that includes a Key/ErrorValue pair
func Error(keyvals ...interface{}) {
	level.Error(logging.logger).Log(keyvals...)
}

// Fatal returns a logger that includes a Key/ErrorValue pair and exit
func Fatal(keyvals ...interface{}) {
	level.Error(logging.logger).Log(keyvals...)
	os.Exit(1)
}
