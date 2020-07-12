package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	. "github.com/yasmindias/travelhelper/models"
)

const Infinity = int(^uint(0) >> 1)

func PopulateGraph(filename string) Graph {
	file := OpenFile(filename)
	routes := ReadFile(file)
	graph := Graph{}

	defer file.Close()

	return graph.AddEdges(routes)
}

func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file
}

func ReadFile(file *os.File) []Route {
	routes := []Route{}

	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		cost, _ := strconv.Atoi(line[2])
		routes = append(routes, Route{line[0], line[1], cost})
	}

	return routes
}

func WriteToFile(file *os.File, route Route) error {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	cost := strconv.Itoa(route.Cost)
	line := []string{route.Origin, route.Destiny, cost}

	err := writer.Write(line)
	return err
}
