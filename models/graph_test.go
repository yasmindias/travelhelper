package models

import (
	"reflect"
	"testing"
)

func setup() *Graph {
	graph := Graph{}
	graph.Init()

	return &graph
}

func addEdges(graph *Graph) {
	graph.AddEdge("S", "B", 4)
	graph.AddEdge("S", "C", 2)
	graph.AddEdge("B", "C", 1)
	graph.AddEdge("B", "D", 5)
	graph.AddEdge("C", "D", 8)
	graph.AddEdge("C", "E", 10)
	graph.AddEdge("D", "E", 2)
	graph.AddEdge("D", "T", 6)
	graph.AddEdge("E", "T", 2)
}

func TestInt(t *testing.T) {
	graph := setup()
	if graph.vertices == nil {
		t.Error("Couldn't create vertices map")
	}
}

func TestDijkstra(t *testing.T) {
	graph := setup()
	addEdges(graph)

	cost, path := graph.Dijkstra("E", "S")
	expCost := 10
	expPath := []string{"E", "D", "B", "C", "S"}

	if cost != expCost {
		t.Error("Couldn't calculate correct cost")
	}

	if !reflect.DeepEqual(path, expPath) {
		t.Error("Couldn't calculate correct path")
	}
}
