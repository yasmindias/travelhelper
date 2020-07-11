package models

import "strconv"

type Path struct {
	value    int
	vertices []string
}

type minPath []Path

func (h minPath) Len() int           { return len(h) }
func (h minPath) Less(i, j int) bool { return h[i].value < h[j].value }
func (h minPath) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(Path))
}

func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (p *Path) ToString() string {
	path := ""
	for i, vertex := range p.vertices {
		path += vertex + " "
		if i != len(p.vertices) {
			path += "- "
		}
	}
	path += " > $" + strconv.Itoa(p.value)
	return path
}
