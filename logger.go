package vlog

// Logger 是最核心的日志接口
type Logger interface {
	Log(level Level, fmt string, args ...any)

	Trace(fmt string, args ...any)
	Debug(fmt string, args ...any)
	Info(fmt string, args ...any)
	Warn(fmt string, args ...any)
	Error(fmt string, args ...any)
	Fatal(fmt string, args ...any)

	IsLevelEnabled(l Level) bool
	IsTraceEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
	IsFatalEnabled() bool

	ForLog(level Level, fn func(logger Logger)) Logger
	ForTrace(fn func(logger Logger)) Logger
	ForDebug(fn func(logger Logger)) Logger
	ForInfo(fn func(logger Logger)) Logger
	ForWarn(fn func(logger Logger)) Logger
	ForError(fn func(logger Logger)) Logger
	ForFatal(fn func(logger Logger)) Logger

	ForSender(sender any) Logger
	ForTag(tag string) Logger
}

// LoggerFactory 是用来创建 Logger 的接口
type LoggerFactory interface {
	Create() Logger
}

// LoggerHolder 是用来获取 Logger 的接口
type LoggerHolder interface {
	Logger() Logger
}
