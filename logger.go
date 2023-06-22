package vlog

// Level 表示日志的等级
type Level int

// 定义日志等级
const (
	ALL   Level = 0 // 把允许的输出等级设置为 ALL，可以输出所有的日志
	TRACE Level = 1
	DEBUG Level = 2
	INFO  Level = 3
	WARN  Level = 4
	ERROR Level = 5
	FATAL Level = 6
	NONE  Level = 99 // 把允许的输出等级设置为 NONE，禁止输出任何等级的日志
)

// Logger 是最核心的日志接口
type Logger interface {
	Trace(fmt string, args ...interface{})
	Debug(fmt string, args ...interface{})
	Info(fmt string, args ...interface{})
	Warn(fmt string, args ...interface{})
	Error(fmt string, args ...interface{})
	Fatal(fmt string, args ...interface{})

	IsTraceEnabled() bool
	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
	IsFatalEnabled() bool
}

// LoggerFactory 是用来创建 Logger 的接口
type LoggerFactory interface {
	Create() Logger
}
