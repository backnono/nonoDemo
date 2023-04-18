package tracing

import (
	"nonoDemo/pkg/framework"
	"fmt"
)

type jaegerLogger struct {
	logger framework.Logger
}

func newLogger(logger framework.Logger) *jaegerLogger {
	return &jaegerLogger{
		logger: logger,
	}
}

func (logger *jaegerLogger) Infof(msg string, args ...interface{}) {
	logger.logger.Info(fmt.Sprintf(msg, args...))
}

func (logger *jaegerLogger) Error(msg string) {
	logger.logger.Error(msg, nil)
}

func (logger *jaegerLogger) Write(p []byte) (int, error) {
	err := logger.logger.Log("msg", string(p))
	return len(p), err
}
