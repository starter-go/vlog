package vlog

import "sort"

////////////////////////////////////////////////////////////////////////////////

// SimpleFilterChainBuilder 用来创建一个 MessageFilterChain
type SimpleFilterChainBuilder struct {
	filters []Filter
}

type MessageFilterChainBuilder = SimpleFilterChainBuilder

// AddFilter 添加一个过滤器
func (inst *SimpleFilterChainBuilder) AddFilter(filter MessageFilter) {
	if filter == nil {
		return
	}
	inst.filters = append(inst.filters, filter)
}

// Create 创建 MessageFilterChain
func (inst *SimpleFilterChainBuilder) Create() FilterChain {
	var chain MessageFilterChain = &myMessageFilterChainEnding{}
	for _, f := range inst.filters {
		if f == nil {
			continue
		}
		node := &myMessageFilterChainNode{filter: f, next: chain}
		chain = node
	}
	return chain
}

////////////////////////////////////////////////////////////////////////////////

type GroupFilterChainBuilder struct {
	items []*FilterRegistration
}

func (inst *GroupFilterChainBuilder) Len() int {
	return len(inst.items)
}
func (inst *GroupFilterChainBuilder) Less(i1, i2 int) bool {
	n1 := inst.items[i1].Order
	n2 := inst.items[i2].Order
	return (n1 > n2)
}
func (inst *GroupFilterChainBuilder) Swap(i1, i2 int) {
	l := inst.items
	l[i1], l[i2] = l[i2], l[i1]
}

func (inst *GroupFilterChainBuilder) sort() {
	sort.Sort(inst)
}

func (inst *GroupFilterChainBuilder) AddRegistry(src ...FilterRegistry) {
	for _, it := range src {
		if it == nil {
			continue
		}
		tmp := it.ListLogFilterRegistration()
		if tmp == nil {
			continue
		}
		inst.AddRegistration(tmp...)
	}
}

func (inst *GroupFilterChainBuilder) AddRegistration(src ...*FilterRegistration) {
	for _, it := range src {
		if inst.innerAcceptRegistration(it) {
			inst.items = append(inst.items, it)
		}
	}
}

func (inst *GroupFilterChainBuilder) innerAcceptRegistration(it *FilterRegistration) bool {

	if it == nil {
		return false
	}

	if !it.Enabled {
		return false
	}

	if it.Filter == nil {
		return false
	}

	return true
}

func (inst *GroupFilterChainBuilder) Build(g Group) FilterChain {
	inst.sort()
	dst := new(SimpleFilterChainBuilder)
	src := inst.items
	for _, it := range src {
		if it.Group == g {
			dst.AddFilter(it.Filter)
		}
	}
	return dst.Create()
}

////////////////////////////////////////////////////////////////////////////////

type myMessageFilterChainNode struct {
	filter MessageFilter
	next   MessageFilterChain
}

func (inst *myMessageFilterChainNode) _Impl() FilterChain {
	return inst
}

func (inst *myMessageFilterChainNode) DoFilter(msg *Message) {
	inst.filter.DoFilter(msg, inst.next)
}

////////////////////////////////////////////////////////////////////////////////

type myMessageFilterChainEnding struct{}

func (inst *myMessageFilterChainEnding) _Impl() FilterChain {
	return inst
}

func (inst *myMessageFilterChainEnding) DoFilter(msg *Message) {
}

////////////////////////////////////////////////////////////////////////////////
