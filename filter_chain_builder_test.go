package vlog

import "testing"

func TestFilterChainBuilder(t *testing.T) {

	g := GroupConsole
	b := new(GroupFilterChainBuilder)
	msg := new(Message)

	// f := new(DateTimeFilter)

	b.AddRegistration(&FilterRegistration{
		Name:    "timer",
		Order:   OrderTime,
		Enabled: true,
		Group:   g,
		Filter:  new(DateTimeFilter),
	})
	b.AddRegistration(&FilterRegistration{
		Name:    "formatter",
		Order:   OrderFormat,
		Enabled: true,
		Group:   g,
		Filter:  new(FormatterFilter),
	})
	b.AddRegistration(&FilterRegistration{
		Order:   OrderWrite,
		Name:    "writer",
		Enabled: true,
		Group:   g,
		Filter:  new(ConsoleDisplayFilter),
	})

	msg.Format = "demo"
	ch := b.Build(g)
	ch.DoFilter(msg)

}
