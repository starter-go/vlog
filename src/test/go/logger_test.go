package main

import (
	"testing"

	"github.com/starter-go/vlog"
)

func TestLogger(t *testing.T) {

	const format = "this is a test of vlog.Logger, index:%v level:%v"
	logger := vlog.GetLogger()

	list := []vlog.Level{
		vlog.TRACE,
		vlog.DEBUG,
		vlog.INFO,
		vlog.WARN,
		vlog.ERROR,
		vlog.FATAL,
	}

	t.Log("step: 1")
	stepFmt1 := "step-1: " + format
	for idx, level := range list {
		logger.ForLog(level, func(l vlog.Logger) {
			l.Log(level, stepFmt1, idx, level.String())
		})
	}

	t.Log("step: 2")
	stepFmt2 := "step-2: " + format
	for idx, level := range list {
		vlog.ForLog(level, func(l vlog.Logger) {
			l.Log(level, stepFmt2, idx, level.String())
		})
	}

	t.Log("done")
}
