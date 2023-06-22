package main

import (
	"errors"

	"github.com/starter-go/vlog"
)

func main() {

	vlog.Trace("this is a trace")
	vlog.Debug("this is a debug")
	vlog.Info("this is a info")
	vlog.Warn("this is a warn")
	vlog.Error("this is a error")
	vlog.Fatal("this is a fatal")

	vlog.Error("error2: ", errors.New(":-("))

	vlog.Info("log int: ", int(-10))
	vlog.Info("log int: ", int8(-80))
	vlog.Info("log int: ", int16(-160))
	vlog.Info("log int: ", int32(-320))
	vlog.Info("log int: ", int64(-640))

	vlog.Info("log int: ", uint(10))
	vlog.Info("log int: ", uint8(80))
	vlog.Info("log int: ", uint16(160))
	vlog.Info("log int: ", uint32(320))
	vlog.Info("log int: ", uint64(640))

	vlog.Info("log float32: ", float32(0.32))
	vlog.Info("log float64: ", float64(0.64))
	vlog.Info("log rune:    ", 'r')
	vlog.Info("log byte:    ", byte(233))
	vlog.Info("log bool:    ", true)
	vlog.Info("log string: ", "str")
}
