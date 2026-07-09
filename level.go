package vlog

import (
	"fmt"
	"strings"
)

// Level 表示日志的等级
type Level int

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
