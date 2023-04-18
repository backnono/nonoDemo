package log

import (
	"context"
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils/observability"
)

type ContextLogger struct {
	logger framework.Logger
	ctx    context.Context
}

func NewContextLogger(ctx context.Context, logger framework.Logger) *ContextLogger {
	return &ContextLogger{
		logger: logger,
		ctx:    ctx,
	}
}

func (logger *ContextLogger) Panic(v ...interface{}) {
	logger.logger.Panic(v, logger.getEventValues())
}

func (logger *ContextLogger) Error(msg string, err error, keyvals ...interface{}) {
	logger.logger.Error(msg, err, append(keyvals, logger.getEventValues()...)...)
}

func (logger *ContextLogger) Warn(msg string, err error, keyvals ...interface{}) {
	logger.logger.Warn(msg, err, append(keyvals, logger.getEventValues()...)...)
}

func (logger *ContextLogger) Info(msg string, keyvals ...interface{}) {
	logger.logger.Info(msg, append(keyvals, logger.getEventValues()...)...)
}

func (logger *ContextLogger) Debug(msg string, keyvals ...interface{}) {
	logger.logger.Debug(msg, append(keyvals, logger.getEventValues()...)...)
}

func (logger *ContextLogger) Log(keyvals ...interface{}) error {
	return logger.logger.Log(append(keyvals, logger.getEventValues()...)...)
}

func (logger *ContextLogger) getEventValues() []interface{} {
	event := observability.GetEvent(logger.ctx)
	return event.EventToArgList()
}
