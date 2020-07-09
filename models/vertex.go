package models

import (
	"fmt"
)

const MAXINT = ^uint(0)

type Vertex struct {
	id         string
	dist       int
	previous   *Vertex
	neighbours map[*Vertex]int
}

func (v *Vertex) PrintPath() {
	if v == v.previous {
		fmt.Print("%s", v.id)
	} else if v.previous == nil {
		fmt.Println("%s is unreachable. ", v.id)
	} else {
		v.previous.PrintPath()
		fmt.Print(" - %s", v.id)
	}
}
