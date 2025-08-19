package vlog

// LoggerBase 提供基本的日志接口
type LoggerBase interface {
	MessageHandler

	IsLevelEnabled(l Level) bool

	Log(level Level, fmt string, args ...any)

	ForLog(level Level, fn func(logger Logger))
}

// 提供扩展的日志接口
type LoggerExt interface {
	LoggerBase

	IsTraceEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
	IsFatalEnabled() bool

	ForTrace(fn func(logger Logger))
	ForDebug(fn func(logger Logger))
	ForInfo(fn func(logger Logger))
	ForWarn(fn func(logger Logger))
	ForError(fn func(logger Logger))
	ForFatal(fn func(logger Logger))

	Trace(fmt string, args ...any)
	Debug(fmt string, args ...any)
	Info(fmt string, args ...any)
	Warn(fmt string, args ...any)
	Error(fmt string, args ...any)
	Fatal(fmt string, args ...any)
}

// Logger 是最核心的日志接口
type Logger interface {
	LoggerExt

	WithSender(sender any) Logger

	WithTag(tag string) Logger
}

// LoggerFactory 是用来创建 Logger 的接口
type LoggerFactory interface {
	Create() Logger
}

// LoggerHolder 是用来获取 Logger 的接口
type LoggerHolder interface {
	Logger() Logger
}
