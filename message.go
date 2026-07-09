package vlog

import "time"

type Group string

type Order int

// Message 表示一条日志记录
type Message struct {
	Arguments []any
	Error     error
	Format    string // 这是用户输入的原始格式, 不包含 message 的头部
	Level     Level
	Tag       string
	Sender    any
	Text      string
	Timestamp time.Time
	Method    string // use: http.Method, 表示需要的操作
	Status    int    // use: http.Status, 表示操作的结果
	Location  string
}

// MessageHandler  表示一个能接收并处理日志消息的接口
type MessageHandler interface {
	HandleMessage(msg *Message)
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
	Name    string
	Enabled bool
	Order   Order
	Group   Group
	Filter  Filter
}

// MessageFilterRegistry 表示一个日志过滤器注册器
type MessageFilterRegistry interface {
	ListLogFilterRegistration() []*MessageFilterRegistration
}

////////////////////////////////////////////////////////////////////////////////
// EOF
