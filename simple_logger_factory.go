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
	builder := MessageFilterChainBuilder{}

	builder.AddFilter(&ConsoleDisplayFilter{})
	builder.AddFilter(&FormatterFilter{})
	builder.AddFilter(&ErrorFilter{})
	builder.AddFilter(&DateTimeFilter{})
	builder.AddFilter(&LevelFilter{AcceptedLevel: level})

	chain := builder.Create()
	return &SimpleLogger{chain: chain, acceptedLevel: level}
}

////////////////////////////////////////////////////////////////////////////////

// SimpleLogger 实现一个简单的日志接口
type SimpleLogger struct {
	chain         MessageFilterChain
	acceptedLevel Level
}

func (inst *SimpleLogger) _Impl() Logger {
	return inst
}

// ForSender 虽然但是 SimpleLogger 不支持 sender
func (inst *SimpleLogger) ForSender(sender any) Logger {
	return inst
}

// ForTag 虽然但是 SimpleLogger 不支持 tag
func (inst *SimpleLogger) ForTag(tag string) Logger {
	return inst
}

func (inst *SimpleLogger) arr(src ...interface{}) []interface{} {
	dst := make([]interface{}, 0)
	for _, item := range src {
		dst = append(dst, item)
	}
	return dst
}

// Trace ...
func (inst *SimpleLogger) Trace(fmt string, args ...interface{}) {
	msg := &Message{Level: TRACE}
	msg.Format = fmt
	msg.Arguments = inst.arr(args...)
	inst.chain.DoFilter(msg)
}

// Debug ...
func (inst *SimpleLogger) Debug(fmt string, args ...interface{}) {
	msg := &Message{Level: DEBUG}
	msg.Format = fmt
	msg.Arguments = inst.arr(args...)
	inst.chain.DoFilter(msg)
}

// Info ...
func (inst *SimpleLogger) Info(fmt string, args ...interface{}) {
	msg := &Message{Level: INFO}
	msg.Format = fmt
	msg.Arguments = inst.arr(args...)
	inst.chain.DoFilter(msg)
}

// Warn ...
func (inst *SimpleLogger) Warn(fmt string, args ...interface{}) {
	msg := &Message{Level: WARN}
	msg.Format = fmt
	msg.Arguments = inst.arr(args...)
	inst.chain.DoFilter(msg)
}

// Error ...
func (inst *SimpleLogger) Error(fmt string, args ...interface{}) {
	msg := &Message{Level: ERROR}
	msg.Format = fmt
	msg.Arguments = inst.arr(args...)
	inst.chain.DoFilter(msg)
}

// Fatal ...
func (inst *SimpleLogger) Fatal(fmt string, args ...interface{}) {
	msg := &Message{Level: FATAL}
	msg.Format = fmt
	msg.Arguments = inst.arr(args...)
	inst.chain.DoFilter(msg)
}

// IsTraceEnabled ...
func (inst *SimpleLogger) IsTraceEnabled() bool {
	return TRACE >= inst.acceptedLevel
}

// IsDebugEnabled ...
func (inst *SimpleLogger) IsDebugEnabled() bool {
	return DEBUG >= inst.acceptedLevel
}

// IsInfoEnabled ...
func (inst *SimpleLogger) IsInfoEnabled() bool {
	return INFO >= inst.acceptedLevel
}

// IsWarnEnabled ...
func (inst *SimpleLogger) IsWarnEnabled() bool {
	return WARN >= inst.acceptedLevel
}

// IsErrorEnabled ...
func (inst *SimpleLogger) IsErrorEnabled() bool {
	return ERROR >= inst.acceptedLevel
}

// IsFatalEnabled ...
func (inst *SimpleLogger) IsFatalEnabled() bool {
	return FATAL >= inst.acceptedLevel
}
