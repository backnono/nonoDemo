package agin

import (
	"nonoDemo/pkg/framework"
)

type ginLoggerAdapter struct {
	logger framework.Logger
}

func newGinLoggerAdapter(logger framework.Logger) *ginLoggerAdapter {
	return &ginLoggerAdapter{logger: logger}
}

func (logger *ginLoggerAdapter) Write(p []byte) (n int, err error) {
	err = logger.logger.Log("msg", string(p))
	return len(p), err
}
