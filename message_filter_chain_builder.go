package vlog

// MessageFilterChainBuilder 用来创建一个 MessageFilterChain
type MessageFilterChainBuilder struct {
	filters []MessageFilter
}

// AddFilter 添加一个过滤器
func (inst *MessageFilterChainBuilder) AddFilter(filter MessageFilter) {
	if filter == nil {
		return
	}
	inst.filters = append(inst.filters, filter)
}

// Create 创建 MessageFilterChain
func (inst *MessageFilterChainBuilder) Create() MessageFilterChain {
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

type myMessageFilterChainNode struct {
	filter MessageFilter
	next   MessageFilterChain
}

func (inst *myMessageFilterChainNode) _Impl() MessageFilterChain {
	return inst
}

func (inst *myMessageFilterChainNode) DoFilter(msg *Message) {
	inst.filter.DoFilter(msg, inst.next)
}

////////////////////////////////////////////////////////////////////////////////

type myMessageFilterChainEnding struct{}

func (inst *myMessageFilterChainEnding) _Impl() MessageFilterChain {
	return inst
}

func (inst *myMessageFilterChainEnding) DoFilter(msg *Message) {
}

////////////////////////////////////////////////////////////////////////////////
