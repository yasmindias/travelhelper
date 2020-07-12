package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	. "github.com/yasmindias/travelhelper/models"
)

const Infinity = int(^uint(0) >> 1)

func PopulateGraph(filename string) Graph {
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

func readFile(file *os.File) Graph {
	graph := Graph{}
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

func writeToFile(filename string, route Route) {
	file := openFile(filename)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	cost := strconv.Itoa(route.Cost)
	line := []string{route.Origin, route.Destiny, cost}

	err := writer.Write(line)
	if err != nil {
		log.Fatal("Can't add new route to file", err)
	}
}
