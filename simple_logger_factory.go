package vlog

// SimpleLoggerFactory 实现一个简易的日志工厂
type SimpleLoggerFactory struct {
	AcceptedLevel Level
}

func (inst *SimpleLoggerFactory) _Impl() LoggerFactory {
	return inst
}

// Create 创建日志接口
func (inst *SimpleLoggerFactory) Create() Logger {

	level := inst.AcceptedLevel

	// chain
	builder := MessageFilterChainBuilder{}

	builder.AddFilter(&ConsoleDisplayFilter{})
	builder.AddFilter(&FormatterFilter{})
	builder.AddFilter(&ErrorFilter{})
	builder.AddFilter(&DateTimeFilter{})
	builder.AddFilter(&LevelFilter{AcceptedLevel: level})

	chain := builder.Create()

	// adapter
	ada := new(LoggerAdapter)

	// bind
	ada.SetTargetChain(chain)
	ada.SetLevelAccepted(level)
	ada.SetSender(inst) // use self
	ada.SetTag("simple-logger")

	return ada
}
