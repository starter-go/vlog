package vlog

var theCurrentLoggerHolder myLoggerHolder
var theDefaultLoggerHolder myLoggerHolder

////////////////////////////////////////////////////////////////////////////////

type myLoggerHolder struct {
	logger  Logger
	factory LoggerFactory
}

func (inst *myLoggerHolder) setFactory(f LoggerFactory) {
	if f == nil {
		return
	}
	inst.factory = f
	inst.logger = nil
}

func (inst *myLoggerHolder) getFactory() LoggerFactory {
	f := inst.factory
	if f == nil {
		slf := new(SimpleLoggerFactory)
		slf.AcceptedLevel = INFO
		f = slf
		inst.factory = f
	}
	return f
}

func (inst *myLoggerHolder) getLogger() Logger {
	l := inst.logger
	if l == nil {
		f := inst.getFactory()
		l = f.Create()
		inst.logger = l
	}
	return l
}

////////////////////////////////////////////////////////////////////////////////

// GetLogger 获取(当前)日志接口
func GetLogger() Logger {
	return theCurrentLoggerHolder.getLogger()
}

// GetLoggerFactory 获取(当前)日志工厂
func GetLoggerFactory() LoggerFactory {
	return theCurrentLoggerHolder.getFactory()
}

// SetLoggerFactory 设置(当前)日志工厂
func SetLoggerFactory(factory LoggerFactory) {
	theCurrentLoggerHolder.setFactory(factory)
}

// GetDefaultLogger 获取(默认)日志接口
func GetDefaultLogger() Logger {
	return theDefaultLoggerHolder.getLogger()
}

// GetDefaultLoggerFactory 获取(默认)日志工厂
func GetDefaultLoggerFactory() LoggerFactory {
	return theDefaultLoggerHolder.getFactory()
}
