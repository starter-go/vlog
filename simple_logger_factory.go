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

////////////////////////////////////////////////////////////////////////////////

// // SimpleLogger 实现一个简单的日志接口
// type SimpleLogger struct {
// 	chain MessageFilterChain

// 	acceptedLevel Level

// 	facade Logger
// }

// func (inst *SimpleLogger) _Impl() LoggerBase {
// 	return inst
// }

// func (inst *SimpleLogger) IsLevelEnabled(l Level) bool {
// 	return l >= inst.acceptedLevel
// }

// func (inst *SimpleLogger) Log(level Level, fmt string, args ...any) {
// 	inst.facade.Log(level, fmt, args...)
// }

// func (inst *SimpleLogger) ForLog(level Level, fn func(logger Logger)) {
// 	inst.facade.ForLog(level, fn)
// }

// func (inst *SimpleLogger) HandleMessage(msg *Message) {
// 	inst.chain.DoFilter(msg)
// }

//  func (inst *SimpleLogger) xxxxxxxx ( ) Logger  {}

// // ForSender 虽然但是 SimpleLogger 不支持 sender
// func (inst *SimpleLogger) ForSender(sender any) Logger {
// 	return inst
// }

// // ForTag 虽然但是 SimpleLogger 不支持 tag
// func (inst *SimpleLogger) ForTag(tag string) Logger {
// 	return inst
// }

// func (inst *SimpleLogger) arr(src ...any) []any {
// 	dst := make([]any, 0)
// 	dst = append(dst, src...)
// 	return dst
// }

// // Trace ...
// func (inst *SimpleLogger) Trace(fmt string, args ...any) {
// 	msg := &Message{Level: TRACE}
// 	msg.Format = fmt
// 	msg.Arguments = inst.arr(args...)
// 	inst.chain.DoFilter(msg)
// }

// // Debug ...
// func (inst *SimpleLogger) Debug(fmt string, args ...any) {
// 	msg := &Message{Level: DEBUG}
// 	msg.Format = fmt
// 	msg.Arguments = inst.arr(args...)
// 	inst.chain.DoFilter(msg)
// }

// // Info ...
// func (inst *SimpleLogger) Info(fmt string, args ...any) {
// 	msg := &Message{Level: INFO}
// 	msg.Format = fmt
// 	msg.Arguments = inst.arr(args...)
// 	inst.chain.DoFilter(msg)
// }

// // Warn ...
// func (inst *SimpleLogger) Warn(fmt string, args ...any) {
// 	msg := &Message{Level: WARN}
// 	msg.Format = fmt
// 	msg.Arguments = inst.arr(args...)
// 	inst.chain.DoFilter(msg)
// }

// // Error ...
// func (inst *SimpleLogger) Error(fmt string, args ...any) {
// 	msg := &Message{Level: ERROR}
// 	msg.Format = fmt
// 	msg.Arguments = inst.arr(args...)
// 	inst.chain.DoFilter(msg)
// }

// // Fatal ...
// func (inst *SimpleLogger) Fatal(fmt string, args ...any) {
// 	msg := &Message{Level: FATAL}
// 	msg.Format = fmt
// 	msg.Arguments = inst.arr(args...)
// 	inst.chain.DoFilter(msg)
// }

// func (inst *SimpleLogger) Log(level Level, fmt string, args ...any) {
// 	msg := &Message{}
// 	msg.Level = level
// 	msg.Format = fmt
// 	msg.Arguments = inst.arr(args...)
// 	inst.chain.DoFilter(msg)
// }

// func (inst *SimpleLogger) IsLevelEnabled(l Level) bool {
// 	return l >= inst.acceptedLevel
// }

// // IsTraceEnabled ...
// func (inst *SimpleLogger) IsTraceEnabled() bool {
// 	return TRACE >= inst.acceptedLevel
// }

// // IsDebugEnabled ...
// func (inst *SimpleLogger) IsDebugEnabled() bool {
// 	return DEBUG >= inst.acceptedLevel
// }

// // IsInfoEnabled ...
// func (inst *SimpleLogger) IsInfoEnabled() bool {
// 	return INFO >= inst.acceptedLevel
// }

// // IsWarnEnabled ...
// func (inst *SimpleLogger) IsWarnEnabled() bool {
// 	return WARN >= inst.acceptedLevel
// }

// // IsErrorEnabled ...
// func (inst *SimpleLogger) IsErrorEnabled() bool {
// 	return ERROR >= inst.acceptedLevel
// }

// // IsFatalEnabled ...
// func (inst *SimpleLogger) IsFatalEnabled() bool {
// 	return FATAL >= inst.acceptedLevel
// }

// func (inst *SimpleLogger) ForFatal(fn func(l Logger)) Logger {
// 	l := inst._Impl()
// 	if inst.isReadyForCallback(fn) {
// 		if inst.IsFatalEnabled() {
// 			fn(l)
// 		}
// 	}
// 	return l
// }

// func (inst *SimpleLogger) ForError(fn func(l Logger)) Logger {
// 	l := inst._Impl()
// 	if inst.isReadyForCallback(fn) {
// 		if inst.IsErrorEnabled() {
// 			fn(l)
// 		}
// 	}
// 	return l
// }
// func (inst *SimpleLogger) ForWarn(fn func(l Logger)) Logger {
// 	l := inst._Impl()
// 	if inst.isReadyForCallback(fn) {
// 		if inst.IsWarnEnabled() {
// 			fn(l)
// 		}
// 	}
// 	return l
// }
// func (inst *SimpleLogger) ForInfo(fn func(l Logger)) Logger {
// 	l := inst._Impl()
// 	if inst.isReadyForCallback(fn) {
// 		if inst.IsInfoEnabled() {
// 			fn(l)
// 		}
// 	}
// 	return l
// }
// func (inst *SimpleLogger) ForDebug(fn func(l Logger)) Logger {
// 	l := inst._Impl()
// 	if inst.isReadyForCallback(fn) {
// 		if inst.IsDebugEnabled() {
// 			fn(l)
// 		}
// 	}
// 	return l
// }

// func (inst *SimpleLogger) ForTrace(fn func(l Logger)) Logger {
// 	l := inst._Impl()
// 	if inst.isReadyForCallback(fn) {
// 		if inst.IsTraceEnabled() {
// 			fn(l)
// 		}
// 	}
// 	return l
// }

// func (inst *SimpleLogger) ForLog(level Level, fn func(logger Logger)) Logger {
// 	l := inst._Impl()
// 	if inst.isReadyForCallback(fn) {
// 		if inst.IsLevelEnabled(level) {
// 			fn(l)
// 		}
// 	}
// 	return l
// }

// func (inst *SimpleLogger) isReadyForCallback(fn func(l Logger)) bool {
// 	if fn == nil {
// 		return false
// 	}
// 	if inst == nil {
// 		return false
// 	}
// 	return true
// }
