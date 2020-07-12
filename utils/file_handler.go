package utils

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	. "github.com/yasmindias/travelhelper/models"
)

func PopulateGraph() Graph {
	file := OpenFileToRead()
	routes := ReadFile(file)
	graph := Graph{}

	defer file.Close()

	return graph.AddEdges(routes)
}

func OpenFileToRead() *os.File {
	file, err := os.Open(os.Getenv("filename"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file
}

func OpenFileToWrite() *os.File {
	file, err := os.OpenFile(os.Getenv("filename"), os.O_APPEND|os.O_WRONLY, 0644)
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

	cost := strconv.Itoa(route.Cost)
	line := []string{route.Origin, route.Destiny, cost}

	err := writer.Write(line)
	writer.Flush()
	file.Close()

	return err
}

func isValidCsvFile(filename string) bool {
	return filename[len(filename)-4:] == ".csv"
}

func StartGraphWithCsvFile(input string) Graph {
	if len(input) > 0 {
		csvFileName := os.Args[1]
		if isValidCsvFile(csvFileName) {
			return PopulateGraph()
		}
	}
	fmt.Println(errors.New("The input must be an existing csv file."))
	os.Exit(1)
	return Graph{}
}
