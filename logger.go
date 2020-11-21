package glog

type Logger interface {
	AddCallerSkip(n int) Logger

	WithField(key string, val interface{}) Logger
	WithFields(key string, val interface{}, kvs ...interface{}) Logger

	Data(val interface{}) Logger
	Action(val string) Logger
	Request(val interface{}) Logger
	Response(val interface{}) Logger

	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}
