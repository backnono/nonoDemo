package log

import (
	"nonoDemo/pkg/framework"
	"github.com/go-kit/log"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelErr
)

type logger struct {
	logger    log.Logger
	logWriter io.Writer
	cfg       Config
	lock      sync.Once
	logLevel  int
	lockMap   sync.Map
}

func NewLogger(cfg Config) framework.Logger {
	// l := log.With(_logger, "pkg", pkg)
	logger := &logger{
		cfg: cfg,
	}
	logger.initLogLevel()
	_logger = logger
	return logger
}

func (logger *logger) Panic(v ...interface{}) {
	panic(v)
}

func (logger *logger) Error(msg string, err error, keyvals ...interface{}) {
	if logger.logLevel <= LevelErr {
		keyvals = append(keyvals, "level", "ERROR", "msg", msg, "err", err)
		_ = logger.write(keyvals...)
	}
}

func (logger *logger) Warn(msg string, err error, keyvals ...interface{}) {
	if logger.logLevel <= LevelWarn {
		keyvals = append(keyvals, "level", "WARN", "msg", msg, "err", err)
		_ = logger.write(keyvals...)
	}
}

func (logger *logger) Info(msg string, keyvals ...interface{}) {
	if logger.logLevel <= LevelInfo {
		keyvals = append(keyvals, "level", "INFO", "msg", msg)
		_ = logger.write(keyvals...)
	}
}

func (logger *logger) Debug(msg string, keyvals ...interface{}) {
	if logger.logLevel <= LevelDebug {
		keyvals = append(keyvals, "level", "DEBUG", "msg", msg)
		_ = logger.write(keyvals...)
	}
}

func (logger *logger) Log(keyvals ...interface{}) error {
	return logger.write(keyvals...)
}

func (logger *logger) Close() {
	logger.logWriter = nil
	logger.logger = nil
}

func (logger *logger) write(keyvals ...interface{}) error {
	logger.getKitLogger()
	return logger.logger.Log(keyvals...)
}

func (logger *logger) getKitLogger() {
	var getLogger func(io.Writer) log.Logger
	getLogger = func(writer io.Writer) log.Logger {
		var kitlogger log.Logger
		if logger.cfg.Format == "text" {
			kitlogger = log.NewLogfmtLogger(logger.logWriter)
		} else if logger.cfg.Format == "json" {
			kitlogger = log.NewJSONLogger(logger.logWriter)
		} else {
			kitlogger = log.NewLogfmtLogger(logger.logWriter)
		}
		kitlogger = log.With(kitlogger, "ts", log.DefaultTimestamp)
		kitlogger = log.With(kitlogger, "caller", Caller(10))
		return kitlogger
	}
	switch logger.cfg.Writer {
	case "os.Stdout":
		logger.lock.Do(func() {
			logger.logWriter = os.Stdout
			logger.logger = getLogger(logger.logWriter)
		})
	case "os.Stderr":
		logger.lock.Do(func() {
			logger.logWriter = os.Stdout
			logger.logger = getLogger(logger.logWriter)
		})
	case "os.File":
		var logPath string
		if logger.cfg.Path[len(logger.cfg.Path)-1] == '/' {
			// user specified path is dir
			name := time.Now().Format("2006-01-02") + ".log"
			var lock interface{} = &sync.Once{}
			lock, _ = logger.lockMap.LoadOrStore(name, lock)
			(lock.(*sync.Once)).Do(func() {
				_ = os.Mkdir(logger.cfg.Path, 0755)
				logPath = path.Join(logger.cfg.Path, name)
				file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					file = os.Stdout
				}
				logger.logWriter = file
				logger.logger = getLogger(logger.logWriter)
			})
		} else {
			logger.lock.Do(func() {
				logPath = logger.cfg.Path
				_ = os.Mkdir(path.Dir(logPath), 0755)
				f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					f = os.Stdout
				}
				logger.logWriter = f
				logger.logger = getLogger(logger.logWriter)
			})

		}

	default:
		logger.lock.Do(func() {
			logger.logWriter = os.Stdout
			logger.logger = getLogger(logger.logWriter)
		})
	}
}

func (logger *logger) initLogLevel() {
	switch logger.cfg.Level {
	case "debug", "DEBUG":
		logger.logLevel = LevelDebug
	case "info", "INFO":
		logger.logLevel = LevelInfo
	case "warn", "WARN":
		logger.logLevel = LevelWarn
	case "err", "ERR":
		logger.logLevel = LevelErr
	default:
		logger.logLevel = LevelInfo
	}
}

var Writer func() io.Writer

var _logger framework.Logger

func GlobalLogger() framework.Logger {
	if _logger != nil {
		return _logger
	}
	return NewLogger(Config{})
}
