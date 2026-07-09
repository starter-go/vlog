package vlog

// 定义日志等级
const (
	MIN Level = 0
	ALL Level = 1 // 把允许的输出等级设置为 ALL，可以输出所有的日志

	TRACE Level = 11
	DEBUG Level = 12
	INFO  Level = 13
	WARN  Level = 14
	ERROR Level = 15
	FATAL Level = 16

	NONE Level = 98 // 把允许的输出等级设置为 NONE，禁止输出任何等级的日志
	MAX  Level = 99
)

const (
	GroupMain    Group = "main"
	GroupConsole Group = "console"
	GroupFile    Group = "file"
	GroupNet     Group = "net"
	GroupDebug   Group = "debug"
)

const (
	OrderLevel  Order = 100
	OrderTime   Order = 200
	OrderFormat Order = 300
	OrderWrite  Order = 400
)
