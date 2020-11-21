package glog

var (
	dflOptions = Options{}
	dflGLog    = New(&dflOptions)
)

func SetDefaultLog(dfl Logger) {
	if l, ok := dfl.(*glog); ok {
		dflGLog = l
	}
}

func GetDefaultLog() Logger { return dflGLog }

func WithField(key string, val interface{}) Logger { return dflGLog.WithField(key, val) }
func WithFields(key string, val interface{}, kvs ...interface{}) Logger {
	return dflGLog.WithFields(key, val, kvs...)
}
func Data(val interface{}) Logger     { return dflGLog.WithField("data", val) }
func Action(val string) Logger        { return dflGLog.WithField("action", val) }
func Request(val interface{}) Logger  { return dflGLog.WithField("request", val) }
func Response(val interface{}) Logger { return dflGLog.WithField("response", val) }
func Debug(msg string)                { dflGLog.AddCallerSkip(1).Debug(msg) }
func Info(msg string)                 { dflGLog.AddCallerSkip(1).Info(msg) }
func Warn(msg string)                 { dflGLog.AddCallerSkip(1).Warn(msg) }
func Error(msg string)                { dflGLog.AddCallerSkip(1).Error(msg) }
func Fatal(msg string)                { dflGLog.AddCallerSkip(1).Fatal(msg) }
