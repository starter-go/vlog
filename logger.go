package vlog

// Logger 是最核心的日志接口
type Logger interface {
	Trace(fmt string, args ...any)
	Debug(fmt string, args ...any)
	Info(fmt string, args ...any)
	Warn(fmt string, args ...any)
	Error(fmt string, args ...any)
	Fatal(fmt string, args ...any)

	IsTraceEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
	IsFatalEnabled() bool

	ForSender(sender any) Logger
	ForTag(tag string) Logger
}

// LoggerFactory 是用来创建 Logger 的接口
type LoggerFactory interface {
	Create() Logger
}
