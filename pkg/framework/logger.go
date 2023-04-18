package framework

type Logger interface {
	Panic(v ...interface{})
	Error(msg string, err error, keyvals ...interface{})
	Warn(msg string, err error, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Debug(msg string, keyvals ...interface{})
	Log(keyvals ...interface{}) error
}
