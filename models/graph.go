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

func (g *Graph) AddEdges(routes []Route) Graph {
	g.Init()
	for _, route := range routes {
		g.addEdge(route.Origin, route.Destiny, route.Cost)
	}
	return *g
}

func (g *Graph) addEdge(origin, destiny string, cost int) {
	g.vertices[origin] = append(g.vertices[origin], Edge{vertex: destiny, cost: cost})
	g.vertices[destiny] = append(g.vertices[destiny], Edge{vertex: origin, cost: cost})
}

func (g *Graph) getEdges(vertex string) []Edge {
	return g.vertices[vertex]
}

func (g *Graph) Dijkstra(origin, destiny string) Path {
	h := newHeap()
	h.push(Path{Cost: 0, Path: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		//Gets the nearest vertex
		p := h.pop()
		vertex := p.Path[len(p.Path)-1]

		if visited[vertex] {
			continue
		}

		if vertex == destiny {
			return p
		}

		for _, e := range g.getEdges(vertex) {
			if !visited[vertex] {
				altCost := p.Cost + e.cost
				vertices := append([]string{}, append(p.Path, e.vertex)...)
				h.push(Path{Cost: altCost, Path: vertices})
			}
		}

		visited[vertex] = true
	}
	return Path{0, nil}
}

func (g *Graph) PrintPath(path Path) string {
	str := ""
	for _, vertex := range path.Path {
		str += vertex + " - "
	}
	str = str[0:len(str)-3] + " > $" + strconv.Itoa(path.Cost)

	return str
}
