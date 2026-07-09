package vlog

import (
	"testing"
	"time"
)

func TestFormatMsgHead(t *testing.T) {

	f := "[{{  .Level }}]  {{.Year}}-{{.Month}}-{{.Day}}  {{.HH}}:{{.MM}}:{{.SS}}.{{.SSS}}  [zone:'{{.zone }}'  tag:'{{.tag}}' sender:'{{.sender}}']"
	m := new(Message)

	m.Timestamp = time.Now()
	m.Level = WARN
	m.Tag = "Demo"
	m.Sender = m

	txt := FormatMessageHead(f, m)

	t.Log("vlog.message.head = ", txt)
}
