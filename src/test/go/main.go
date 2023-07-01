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

	vlog.Error("error2: %v", errors.New(":-("))

	vlog.Info("log int: %d", int(-10))
	vlog.Info("log int: %d", int8(-80))
	vlog.Info("log int: %d", int16(-160))
	vlog.Info("log int: %d", int32(-320))
	vlog.Info("log int: %d", int64(-640))

	vlog.Info("log int: %d", uint(10))
	vlog.Info("log int: %d", uint8(80))
	vlog.Info("log int: %d", uint16(160))
	vlog.Info("log int: %d", uint32(320))
	vlog.Info("log int: %d", uint64(640))

	vlog.Info("log float32: %v", float32(0.32))
	vlog.Info("log float64: %v", float64(0.64))
	vlog.Info("log rune:    %d", 'r')
	vlog.Info("log byte:    %d", byte(233))
	vlog.Info("log bool:    %v", true)
	vlog.Info("log string: %s", "str")
}
