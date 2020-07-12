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

func AddEdges(graph *Graph) {
	routes := []Route{
		Route{"S", "B", 4},
		Route{"S", "C", 2},
		Route{"B", "C", 1},
		Route{"B", "D", 5},
		Route{"C", "D", 8},
		Route{"C", "E", 10},
		Route{"D", "E", 2},
		Route{"D", "T", 6},
		Route{"E", "T", 2},
	}
	graph.AddEdges(routes)
}

func TestInt(t *testing.T) {
	graph := setup()
	if graph.vertices == nil {
		t.Error("Couldn't create vertices map")
	}
}

func TestDijkstra(t *testing.T) {
	graph := setup()
	AddEdges(graph)

	result := graph.Dijkstra("E", "S")
	expCost := 10
	expPath := []string{"E", "D", "B", "C", "S"}

	if result.Cost != expCost {
		t.Error("Couldn't calculate correct cost")
	}

	if !reflect.DeepEqual(result.Path, expPath) {
		t.Error("Couldn't calculate correct path")
	}
}
