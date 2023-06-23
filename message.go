package vlog

import "time"

// Message 表示一条日志记录
type Message struct {
	Arguments []any
	Error     error
	Level     Level
	Text      string
	Timestamp time.Time
}

// MessageFilter 表示一个日志消息过滤器
type MessageFilter interface {
	DoFilter(msg *Message, chain MessageFilterChain)
}

// MessageFilterChain 表示一个日志消息过滤器串
type MessageFilterChain interface {
	DoFilter(msg *Message)
}

// Stringer 表示一个可字符化的接口
type Stringer interface {
	String() string
}
