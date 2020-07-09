package models

import (
	"errors"
	"fmt"
)

const MaxInt = ^int(0)

type Graph struct {
	graph map[string]Vertex
}

func (g *Graph) Create(edges []Edge) {
	g.graph = map[string]Vertex{}

	for _, edge := range edges {
		if _, notExists := g.graph[edge.v1]; notExists {
			g.graph[edge.v1] = Vertex{edge.v1, MaxInt, nil, map[*Vertex]int{}}
		}
		if _, notExists := g.graph[edge.v2]; notExists {
			g.graph[edge.v2] = Vertex{edge.v2, MaxInt, nil, map[*Vertex]int{}}
		}
	}

	for _, edge := range edges {
		vertex1 := g.graph[edge.v1]
		vertex2 := g.graph[edge.v2]
		vertex1.neighbours[&vertex2] = edge.dist
	}
}

func (g *Graph) Dijkstra(startVertex string) {
	if _, notExists := g.graph[startVertex]; notExists {
		fmt.Println(errors.New("Graph doesn't contain vertex '" + startVertex + "'"))
		return
	}

	source := g.graph[startVertex]
	var q map[string]Vertex

	for _, vertex := range g.graph {
		if vertex.id == source.id {
			vertex.previous = &source
			vertex.dist = 0
		} else {
			vertex.previous = nil
			vertex.dist = MaxInt
		}

		q[vertex.id] = vertex
	}

	dijkstra(q)
}

func dijkstra(q map[string]Vertex) {
	var u Vertex

	for key, vertex := range q {
		u = vertex //gets vertex with the smallest distance
		delete(q, key)
		if u.dist == MaxInt {
			break
		}

		for neighbour, dist := range u.neighbours {
			altDistance := u.dist + dist
			if altDistance < dist {
				delete(q, neighbour.id)
				neighbour.dist = altDistance
				neighbour.previous = &u
				q[neighbour.id] = *neighbour
			}
		}
	}
}

func (g *Graph) PrintPath(endVertex string) {
	if _, notExists := g.graph[endVertex]; notExists {
		fmt.Println(errors.New("Graph doesn't contain vertex '" + endVertex + "'"))
		return
	}

	end := g.graph[endVertex]
	end.PrintPath()

	if end.dist != MaxInt {
		fmt.Printf(" > %d \n", end.dist)
	}
}
