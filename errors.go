package vlog

import "fmt"

// HandleError 把错误信息(error)通过日志输出
func HandleError(err error) {
	if err == nil {
		return
	}
	Error("%s", err.Error())
}

// HandleErrorX 把错误信息(any)通过日志输出
func HandleErrorX(x any) {
	if x == nil {
		return
	}
	err := x2error(x)
	HandleError(err)
}

// x2error 把 x 转换成 error
func x2error(x any) error {

	err, ok := x.(error)
	if ok {
		return err
	}

	str, ok := x.(string)
	if ok {
		return fmt.Errorf("%s", str)
	}

	return fmt.Errorf("%s", x)
}
