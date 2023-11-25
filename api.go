package vlog

// Trace 输出一条等级为 [TRACE] 的日志记录
func Trace(fmt string, args ...any) {
	GetLogger().Trace(fmt, args...)
}

// Debug 输出一条等级为 [DEBUG] 的日志记录
func Debug(fmt string, args ...any) {
	GetLogger().Debug(fmt, args...)
}

// Info 输出一条等级为 [INFO] 的日志记录
func Info(fmt string, args ...any) {
	GetLogger().Info(fmt, args...)
}

// Warn 输出一条等级为 [WARN] 的日志记录
func Warn(fmt string, args ...any) {
	GetLogger().Warn(fmt, args...)
}

// Error 输出一条等级为 [ERROR] 的日志记录
func Error(fmt string, args ...any) {
	GetLogger().Error(fmt, args...)
}

// Fatal 输出一条等级为 [FATAL] 的日志记录
func Fatal(fmt string, args ...any) {
	GetLogger().Fatal(fmt, args...)
}

// IsTraceEnabled 判断是否需要输出等级为 [TRACE] 的日志记录
func IsTraceEnabled() bool {
	return GetLogger().IsTraceEnabled()
}

// IsDebugEnabled 判断是否需要输出等级为 [DEBUG] 的日志记录
func IsDebugEnabled() bool {
	return GetLogger().IsDebugEnabled()
}

// IsInfoEnabled 判断是否需要输出等级为 [INFO] 的日志记录
func IsInfoEnabled() bool {
	return GetLogger().IsInfoEnabled()
}

// IsWarnEnabled 判断是否需要输出等级为 [WARN] 的日志记录
func IsWarnEnabled() bool {
	return GetLogger().IsWarnEnabled()
}

// IsErrorEnabled 判断是否需要输出等级为 [ERROR] 的日志记录
func IsErrorEnabled() bool {
	return GetLogger().IsErrorEnabled()
}

// IsFatalEnabled 判断是否需要输出等级为 [FATAL] 的日志记录
func IsFatalEnabled() bool {
	return GetLogger().IsFatalEnabled()
}
