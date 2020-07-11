package models

import hp "container/heap"

type heap struct {
	values *minPath
}

func newHeap() *heap {
	return &heap{values: &minPath{}}
}

func (h *heap) push(p Path) {
	hp.Push(h.values, p)
}

func (h *heap) pop() Path {
	i := hp.Pop(h.values)
	return i.(Path)
}
