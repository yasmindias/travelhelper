package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/yasmindias/travelhelper/models"
)

const Infinity = int(^uint(0) >> 1)

func PopulateGraph(filename string) models.Graph {
	file := openFile(filename)
	graph := readFile(file)

	defer file.Close()

	return graph
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file
}

func readFile(file *os.File) models.Graph {
	graph := models.Graph{}
	graph.Init()

	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		cost, _ := strconv.Atoi(line[2])
		graph.AddEdge(line[0], line[1], cost)
	}

	return graph
}
