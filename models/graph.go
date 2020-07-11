package models

import (
	"strconv"
)

const Infinity = int(^uint(0) >> 1)

type Graph struct {
	vertices map[string][]Edge
}

func (g *Graph) Init() {
	g.vertices = make(map[string][]Edge)
}

func (g *Graph) AddEdge(origin, destiny string, cost int) {
	g.vertices[origin] = append(g.vertices[origin], Edge{vertex: destiny, cost: cost})
	g.vertices[destiny] = append(g.vertices[destiny], Edge{vertex: origin, cost: cost})
}

func (g *Graph) getEdges(vertex string) []Edge {
	return g.vertices[vertex]
}

func (g *Graph) Dijkstra(origin, destiny string) (int, []string) {
	h := newHeap()
	h.push(Path{value: 0, vertices: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		//Gets the nearest vertex
		p := h.pop()
		vertex := p.vertices[len(p.vertices)-1]

		if visited[vertex] {
			continue
		}

		if vertex == destiny {
			return p.value, p.vertices
		}

		for _, e := range g.getEdges(vertex) {
			if !visited[vertex] {
				altCost := p.value + e.cost
				vertices := append([]string{}, append(p.vertices, e.vertex)...)
				h.push(Path{value: altCost, vertices: vertices})
			}
		}

		visited[vertex] = true
	}
	return 0, nil
}

func (g *Graph) PrintPath(cost int, path []string) string {
	str := ""
	for _, vertex := range path {
		str += vertex + " - "
	}
	str = str[0:len(str)-3] + " > $" + strconv.Itoa(cost)

	return str
}
