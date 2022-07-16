package vlog

import (
	"fmt"
	"strings"
	"time"
)

////////////////////////////////////////////////////////////////////////////////

// ErrorFilter  过滤器
type ErrorFilter struct {
}

func (inst *ErrorFilter) _Impl() MessageFilter {
	return inst
}

// DoFilter 执行过滤操作
func (inst *ErrorFilter) DoFilter(msg *Message, chain MessageFilterChain) {
	msg.Error = inst.findError(msg)
	chain.DoFilter(msg)
}

func (inst *ErrorFilter) findError(msg *Message) error {
	for _, o := range msg.Arguments {
		err, ok := o.(error)
		if ok {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// LevelFilter  过滤器
type LevelFilter struct {
	AcceptedLevel Level
}

func (inst *LevelFilter) _Impl() MessageFilter {
	return inst
}

// DoFilter 执行过滤操作
func (inst *LevelFilter) DoFilter(msg *Message, chain MessageFilterChain) {
	if msg.Level >= inst.AcceptedLevel {
		chain.DoFilter(msg)
	}
}

////////////////////////////////////////////////////////////////////////////////

// FormatterFilter  过滤器
type FormatterFilter struct {
}

func (inst *FormatterFilter) _Impl() MessageFilter {
	return inst
}

// DoFilter 执行过滤操作
func (inst *FormatterFilter) DoFilter(msg *Message, chain MessageFilterChain) {
	msg.Text = inst.format(msg)
	chain.DoFilter(msg)
}

func (inst *FormatterFilter) format(msg *Message) string {
	builder := strings.Builder{}
	inst.formatTime(msg, &builder)
	inst.formatLevel(msg, &builder)
	inst.formatArgs(msg.Arguments, &builder)
	return builder.String()
}

func (inst *FormatterFilter) formatLevel(msg *Message, builder *strings.Builder) {
	builder.WriteString(" [")
	switch msg.Level {
	case INFO:
		builder.WriteString("INFO ")
		break
	case DEBUG:
		builder.WriteString("DEBUG")
		break
	case TRACE:
		builder.WriteString("TRACE")
		break
	case WARN:
		builder.WriteString("WARN ")
		break
	case ERROR:
		builder.WriteString("ERROR")
		break
	case FATAL:
		builder.WriteString("FATAL")
		break
	default:
		builder.WriteString("UNDEF")
		break
	}
	builder.WriteString("] ")
}

func (inst *FormatterFilter) formatTime(msg *Message, builder *strings.Builder) {
	t := msg.Timestamp
	str := t.Format(time.RFC3339)
	builder.WriteString(str)
}

func (inst *FormatterFilter) formatArgs(args []interface{}, builder *strings.Builder) {
	for _, a := range args {
		str1, ok := a.(string)
		if ok {
			builder.WriteString(str1)
			continue
		}

		str2, ok := a.(Stringer)
		if ok {
			builder.WriteString(str2.String())
			continue
		}

		err, ok := a.(error)
		if ok {
			builder.WriteString(err.Error())
			continue
		}

		str3 := fmt.Sprint(a)
		builder.WriteString(str3)
	}
}

////////////////////////////////////////////////////////////////////////////////

// DateTimeFilter  过滤器
type DateTimeFilter struct {
}

func (inst *DateTimeFilter) _Impl() MessageFilter {
	return inst
}

// DoFilter 执行过滤操作
func (inst *DateTimeFilter) DoFilter(msg *Message, chain MessageFilterChain) {
	msg.Timestamp = time.Now()
	chain.DoFilter(msg)
}

////////////////////////////////////////////////////////////////////////////////

// ConsoleDisplayFilter  过滤器
type ConsoleDisplayFilter struct {
}

func (inst *ConsoleDisplayFilter) _Impl() MessageFilter {
	return inst
}

// DoFilter 执行过滤操作
func (inst *ConsoleDisplayFilter) DoFilter(msg *Message, chain MessageFilterChain) {
	text := msg.Text
	fmt.Println(text)
	chain.DoFilter(msg)
}

////////////////////////////////////////////////////////////////////////////////
