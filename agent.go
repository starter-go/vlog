package vlog

var theLogger Logger
var theLoggerFactory LoggerFactory

// GetLogger 获取日志接口
func GetLogger() Logger {
	logger := theLogger
	if logger == nil {
		factory := GetLoggerFactory()
		logger = factory.Create()
		theLogger = logger
	}
	return logger
}

// GetLoggerFactory 获取日志工厂
func GetLoggerFactory() LoggerFactory {
	factory := theLoggerFactory
	if factory == nil {
		factory = &SimpleLoggerFactory{AcceptedLevel: INFO}
		theLoggerFactory = factory
	}
	return factory
}

// SetLoggerFactory 设置日志工厂
func SetLoggerFactory(factory LoggerFactory) {
	if factory == nil {
		return
	}
	theLoggerFactory = factory
	theLogger = nil
}
