package vlog

import "time"

type innerLoggerAdapter struct {
	sender        any
	tag           string
	target        MessageHandler
	acceptedLevel Level
}

////////////////////////////////////////////////////////////////////////////////

type innerLoggerChainAdapter struct {
	chain MessageFilterChain
}

func (inst *innerLoggerChainAdapter) _impl() MessageHandler {
	return inst
}

func (inst *innerLoggerChainAdapter) HandleMessage(msg *Message) {
	inst.chain.DoFilter(msg)
}

////////////////////////////////////////////////////////////////////////////////

type LoggerAdapter struct {
	innerLoggerAdapter
}

func (inst *LoggerAdapter) _impl() Logger {
	return inst
}

func (inst *LoggerAdapter) SetTargetChain(chain MessageFilterChain) *LoggerAdapter {
	ada2 := new(innerLoggerChainAdapter)
	ada2.chain = chain
	inst.target = ada2
	return inst
}

func (inst *LoggerAdapter) SetTargetHandler(h MessageHandler) *LoggerAdapter {
	inst.target = h
	return inst
}

func (inst *LoggerAdapter) SetTargetLogger(l LoggerBase) *LoggerAdapter {
	inst.target = l
	return inst
}

func (inst *LoggerAdapter) SetLevelAccepted(lv Level) *LoggerAdapter {
	inst.acceptedLevel = lv
	return inst
}

func (inst *LoggerAdapter) SetTag(tag string) *LoggerAdapter {
	inst.tag = tag
	return inst
}

func (inst *LoggerAdapter) SetSender(sender any) *LoggerAdapter {
	inst.sender = sender
	return inst
}

func (inst *LoggerAdapter) innerCloneMyself() *LoggerAdapter {
	o2 := new(LoggerAdapter)
	o2.innerLoggerAdapter = inst.innerLoggerAdapter
	return o2
}

func (inst *LoggerAdapter) IsLevelEnabled(l Level) bool {
	return l >= inst.acceptedLevel
}

func (inst *LoggerAdapter) ForLog(level Level, fn func(logger Logger)) {

	if fn == nil || inst == nil {
		return
	}

	if !inst.IsLevelEnabled(level) {
		return
	}

	fn(inst)
}

func (inst *LoggerAdapter) WithSender(sender any) Logger {
	o2 := inst.innerCloneMyself()
	o2.sender = sender
	return o2
}

func (inst *LoggerAdapter) WithTag(tag string) Logger {
	o2 := inst.innerCloneMyself()
	o2.tag = tag
	return o2
}

func (inst *LoggerAdapter) HandleMessage(msg *Message) {
	inst.target.HandleMessage(msg)
}

func (inst *LoggerAdapter) Log(level Level, fmt string, args ...any) {

	msg := new(Message)

	msg.Level = level
	msg.Format = fmt
	msg.Arguments = args
	msg.Timestamp = time.Now()
	msg.Sender = inst.sender
	msg.Tag = inst.tag
	msg.Text = ""
	msg.Error = nil

	inst.target.HandleMessage(msg)
}

func (inst *LoggerAdapter) Trace(fmt string, args ...any) {
	const lv = TRACE
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerAdapter) Warn(fmt string, args ...any) {
	const lv = WARN
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerAdapter) Debug(fmt string, args ...any) {
	const lv = DEBUG
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerAdapter) Error(fmt string, args ...any) {
	const lv = ERROR
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerAdapter) Fatal(fmt string, args ...any) {
	const lv = FATAL
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerAdapter) Info(fmt string, args ...any) {
	const lv = INFO
	inst.Log(lv, fmt, args...)
}

func (inst *LoggerAdapter) IsDebugEnabled() bool {
	const lv = DEBUG
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerAdapter) IsErrorEnabled() bool {
	const lv = ERROR
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerAdapter) IsFatalEnabled() bool {
	const lv = FATAL
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerAdapter) IsInfoEnabled() bool {
	const lv = INFO
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerAdapter) IsTraceEnabled() bool {
	const lv = TRACE
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerAdapter) IsWarnEnabled() bool {
	const lv = WARN
	return inst.IsLevelEnabled(lv)
}

func (inst *LoggerAdapter) ForDebug(fn func(logger Logger)) {
	const lv = DEBUG
	inst.ForLog(lv, fn)
}

func (inst *LoggerAdapter) ForError(fn func(logger Logger)) {
	const lv = ERROR
	inst.ForLog(lv, fn)
}

func (inst *LoggerAdapter) ForFatal(fn func(logger Logger)) {
	const lv = FATAL
	inst.ForLog(lv, fn)
}

func (inst *LoggerAdapter) ForInfo(fn func(logger Logger)) {
	const lv = INFO
	inst.ForLog(lv, fn)
}

func (inst *LoggerAdapter) ForTrace(fn func(logger Logger)) {
	const lv = TRACE
	inst.ForLog(lv, fn)
}

func (inst *LoggerAdapter) ForWarn(fn func(logger Logger)) {
	const lv = WARN
	inst.ForLog(lv, fn)
}
