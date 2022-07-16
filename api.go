package vlog

// Trace 输出一条等级为 [TRACE] 的日志记录
func Trace(args ...interface{}) {
	GetLogger().Trace(args...)
}

// Debug 输出一条等级为 [DEBUG] 的日志记录
func Debug(args ...interface{}) {
	GetLogger().Debug(args...)
}

// Info 输出一条等级为 [INFO] 的日志记录
func Info(args ...interface{}) {
	GetLogger().Info(args...)
}

// Warn 输出一条等级为 [WARN] 的日志记录
func Warn(args ...interface{}) {
	GetLogger().Warn(args...)
}

// Error 输出一条等级为 [ERROR] 的日志记录
func Error(args ...interface{}) {
	GetLogger().Error(args...)
}

// Fatal 输出一条等级为 [FATAL] 的日志记录
func Fatal(args ...interface{}) {
	GetLogger().Fatal(args...)
}

// IsTraceEnabled 判断是否需要输出等级为 [TRACE] 的日志记录
func IsTraceEnabled() bool {
	return GetLogger().IsTraceEnabled()
}

// IsDebugEnabled 判断是否需要输出等级为 [DEBUG] 的日志记录
func IsDebugEnabled() bool {
	return GetLogger().IsTraceEnabled()
}

// IsInfoEnabled 判断是否需要输出等级为 [INFO] 的日志记录
func IsInfoEnabled() bool {
	return GetLogger().IsTraceEnabled()
}

// IsWarnEnabled 判断是否需要输出等级为 [WARN] 的日志记录
func IsWarnEnabled() bool {
	return GetLogger().IsTraceEnabled()
}

// IsErrorEnabled 判断是否需要输出等级为 [ERROR] 的日志记录
func IsErrorEnabled() bool {
	return GetLogger().IsTraceEnabled()
}

// IsFatalEnabled 判断是否需要输出等级为 [FATAL] 的日志记录
func IsFatalEnabled() bool {
	return GetLogger().IsTraceEnabled()
}
