package vlog

import "time"

// Message 表示一条日志记录
type Message struct {
	Arguments []any
	Error     error
	Format    string
	Level     Level
	Tag       string
	Sender    any
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

// MessageFilterGroup 表示一个日志消息过滤器组
type MessageFilterGroup interface {
	GetFilterChain() MessageFilterChain
}

// Stringer 表示一个可字符化的接口
type Stringer interface {
	String() string
}

// MessageFilterRegistration 表示一个日志过滤器注册信息
type MessageFilterRegistration struct {
	Order  int
	Name   string
	Filter MessageFilter
}

// MessageFilterRegistry 表示一个日志过滤器注册器
type MessageFilterRegistry interface {
	ListLogFilterRegistration() []*MessageFilterRegistration
}
