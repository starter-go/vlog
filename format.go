package vlog

import (
	"reflect"
	"strconv"
	"strings"
)

func FormatMessageHead(format string, msg *Message) string {
	f := new(MessageHeadFormatter)
	f.Init(format)
	return f.Format(msg)
}

////////////////////////////////////////////////////////////////////////////////

type MessageHeadFormatter struct {
	format string
	parts  []*innerMessageHeadFormattingPart
}

func (inst *MessageHeadFormatter) Init(format string) {
	const (
		mark1 = "{{"
		mark2 = "}}"
	)
	dst := make([]*innerMessageHeadFormattingPart, 0)
	src := strings.Split(format, mark1)
	for _, element := range src {
		ab := strings.SplitN(element, mark2, 2)
		if len(ab) == 2 {
			a := ab[0]
			b := ab[1]
			p0 := new(innerMessageHeadFormattingPart)
			p1 := new(innerMessageHeadFormattingPart)
			p0.initAsKV(a)
			p1.initAsText(b)
			dst = append(dst, p0)
			dst = append(dst, p1)
		} else {
			pn := new(innerMessageHeadFormattingPart)
			pn.initAsText(element)
			dst = append(dst, pn)
		}
	}
	inst.parts = dst
}

func (inst *MessageHeadFormatter) Format(msg *Message) string {

	if msg == nil {
		return ""
	}

	c := new(innerMessageHeadFormattingContext)
	b := new(strings.Builder)

	c.initWithMessage(msg)
	c.formatter = inst
	c.writeTo(b)

	return b.String()
}

////////////////////////////////////////////////////////////////////////////////

type innerMessageHeadFormattingContext struct {
	formatter *MessageHeadFormatter
	fields    map[string]string
}

func (inst *innerMessageHeadFormattingContext) initWithMessage(msg *Message) {

	const (
		fnameYear  = "YYYY"
		fnameMonth = "MM"
		fnameDay   = "DD"
		fnameHour  = "H"
		fnameMin   = "M"
		fnameSec   = "S"
		fnameMS    = "SSS"

		fnameTag    = "TAG"
		fnameLevel  = "LEVEL"
		fnameSender = "SENDER"
	)

	table := inst.fields
	table = make(map[string]string)

	ts := msg.Timestamp
	yyyy := ts.Year()
	mm := int(ts.Month())
	dd := ts.Day()
	h := ts.Hour()
	m := ts.Minute()
	s := ts.Second()
	sss := int(ts.UnixMilli() % 1000)

	table[fnameYear] = inst.innerFormatInt(yyyy, 0)
	table[fnameMonth] = inst.innerFormatInt(mm, 2)
	table[fnameDay] = inst.innerFormatInt(dd, 2)
	table[fnameHour] = inst.innerFormatInt(h, 2)
	table[fnameMin] = inst.innerFormatInt(m, 2)
	table[fnameSec] = inst.innerFormatInt(s, 2)
	table[fnameMS] = inst.innerFormatInt(sss, 3)

	table[fnameTag] = msg.Tag
	table[fnameLevel] = inst.innerFormatLevel(msg.Level)
	table[fnameSender] = inst.innerFormatSender(msg)

	inst.fields = table
}

func (inst *innerMessageHeadFormattingContext) innerFormatSender(msg *Message) string {
	s := msg.Sender
	if s == nil {
		return ""
	}
	t := reflect.TypeOf(s)
	txt := t.String()
	return txt
}

func (inst *innerMessageHeadFormattingContext) innerFormatLevel(lv Level) string {
	const width = 5
	txt := lv.String()
	ttl := 99
	for ; ttl > 0; ttl-- {
		if len(txt) < width {
			txt = "." + txt
		} else {
			break
		}
	}
	return txt
}

func (inst *innerMessageHeadFormattingContext) innerFormatInt(n int, width int) string {
	ttl := 99
	txt := strconv.Itoa(n)
	for ; ttl > 0; ttl-- {
		if len(txt) < width {
			txt = "0" + txt
		} else {
			break
		}
	}
	return txt
}

func (inst *innerMessageHeadFormattingContext) writeTo(b *strings.Builder) {

	table := inst.fields
	parts := inst.formatter.parts

	for _, p := range parts {
		if p.asKeyValue {
			name := p.name
			value := table[name]
			if value == "" {
				value = p.value
			}
			b.WriteString(value)
		} else {
			b.WriteString(p.value)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

type innerMessageHeadFormattingPart struct {
	value      string
	name       string
	asKeyValue bool
}

func (inst *innerMessageHeadFormattingPart) initAsKV(str string) {
	const (
		m1 = "{{"
		m2 = "}}"
	)
	str = strings.TrimSpace(str)
	str = strings.ToUpper(str)
	inst.name = str
	inst.value = m1 + str + m2
	inst.asKeyValue = true
}

func (inst *innerMessageHeadFormattingPart) initAsText(str string) {
	inst.name = ""
	inst.value = str
	inst.asKeyValue = false
}

////////////////////////////////////////////////////////////////////////////////
