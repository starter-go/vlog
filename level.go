package vlog

import (
	"fmt"
	"strings"
)

// Level 表示日志的等级
type Level int

// 定义日志等级
const (
	ALL   Level = 0 // 把允许的输出等级设置为 ALL，可以输出所有的日志
	TRACE Level = 1
	DEBUG Level = 2
	INFO  Level = 3
	WARN  Level = 4
	ERROR Level = 5
	FATAL Level = 6
	NONE  Level = 99 // 把允许的输出等级设置为 NONE，禁止输出任何等级的日志
)

// ParseLevel 把字符串解析为 vlog.Level 值
func ParseLevel(str string) (Level, error) {
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	switch str {
	case "ALL":
		return ALL, nil
	case "TRACE":
		return TRACE, nil
	case "DEBUG":
		return DEBUG, nil
	case "INFO":
		return INFO, nil
	case "WARN":
		return WARN, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	case "NONE":
		return NONE, nil
	}
	return 0, fmt.Errorf("bad vlog.Level string: '%s'", str)
}

func (l Level) String() string {
	switch l {
	case ALL:
		return "ALL"
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "NONE"
}
