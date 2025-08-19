package vlog

import "time"

type innerLoggerContext struct {

	// handler MessageHandler

	tag    string
	inner  LoggerBase
	sender any
}

////////////////////////////////////////////////////////////////////////////////

type LoggerWrapper struct {
	innerLoggerContext
}

func (inst *LoggerWrapper) _impl() Logger {
	return inst
}

func (inst *LoggerWrapper) innerCloneMyself() *LoggerWrapper {
	o2 := new(LoggerWrapper)
	o2.innerLoggerContext = inst.innerLoggerContext
	return o2
}

func (inst *LoggerWrapper) IsLevelEnabled(l Level) bool {
	return inst.inner.IsLevelEnabled(l)
}

func (inst *LoggerWrapper) ForLog(level Level, fn func(logger Logger)) {

	if fn == nil || inst == nil {
		return
	}

	if !inst.IsLevelEnabled(level) {
		return
	}

	fn(inst)
}

func (inst *LoggerWrapper) WithSender(sender any) Logger {
	o2 := inst.innerCloneMyself()
	o2.sender = sender
	return o2
}

func (inst *LoggerWrapper) WithTag(tag string) Logger {
	o2 := inst.innerCloneMyself()
	o2.tag = tag
	return o2
}

func (inst *LoggerWrapper) HandleMessage(msg *Message) {
	inst.inner.HandleMessage(msg)
}

func (inst *LoggerWrapper) Log(level Level, fmt string, args ...any) {

	msg := new(Message)

	msg.Level = level
	msg.Format = fmt
	msg.Arguments = args
	msg.Timestamp = time.Now()
	msg.Sender = inst.sender
	msg.Tag = inst.tag
	msg.Text = ""
	msg.Error = nil

	inst.inner.HandleMessage(msg)
}

func (inst *LoggerWrapper) Trace(fmt string, args ...any) {
	const lv = TRACE
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerWrapper) Warn(fmt string, args ...any) {
	const lv = WARN
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerWrapper) Debug(fmt string, args ...any) {
	const lv = DEBUG
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerWrapper) Error(fmt string, args ...any) {
	const lv = ERROR
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerWrapper) Fatal(fmt string, args ...any) {
	const lv = FATAL
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerWrapper) Info(fmt string, args ...any) {
	const lv = INFO
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerWrapper) IsDebugEnabled() bool {
	const lv = DEBUG
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerWrapper) IsErrorEnabled() bool {
	const lv = ERROR
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerWrapper) IsFatalEnabled() bool {
	const lv = FATAL
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerWrapper) IsInfoEnabled() bool {
	const lv = INFO
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerWrapper) IsTraceEnabled() bool {
	const lv = TRACE
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerWrapper) IsWarnEnabled() bool {
	const lv = WARN
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerWrapper) ForDebug(fn func(logger Logger)) {
	const lv = DEBUG
	inst.ForLog(lv, fn)
}

func (inst *LoggerWrapper) ForError(fn func(logger Logger)) {
	const lv = ERROR
	inst.ForLog(lv, fn)
}

func (inst *LoggerWrapper) ForFatal(fn func(logger Logger)) {
	const lv = FATAL
	inst.ForLog(lv, fn)
}

func (inst *LoggerWrapper) ForInfo(fn func(logger Logger)) {
	const lv = INFO
	inst.ForLog(lv, fn)
}

func (inst *LoggerWrapper) ForTrace(fn func(logger Logger)) {
	const lv = TRACE
	inst.ForLog(lv, fn)
}

func (inst *LoggerWrapper) ForWarn(fn func(logger Logger)) {
	const lv = WARN
	inst.ForLog(lv, fn)
}
