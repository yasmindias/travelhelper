package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/yasmindias/travelhelper/models"
	"github.com/yasmindias/travelhelper/utils"
)

func main() {
	graph := startGraphWithCsvFile(os.Args[1])

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter the route: ")
	route, _ := reader.ReadString('\n')
	if isValidRoute(route) {
		places := strings.Split(strings.ToUpper(route), "-")
		origin := strings.Trim(places[0], " ")
		destiny := strings.TrimSuffix(places[1], "\n")

		cost, path := graph.Dijkstra(origin, destiny)
		fmt.Print(path)
		fmt.Printf(" > $%d\n", cost)
	} else {
		fmt.Println(errors.New("The input must be in the format \"ORG-DEST\"."))
	}
}

func isValidCsvFile(filename string) bool {
	return filename[len(filename)-4:] == ".csv"
}

func startGraphWithCsvFile(input string) models.Graph {
	if len(input) > 0 {
		csvFileName := os.Args[1]
		if isValidCsvFile(csvFileName) {
			return utils.PopulateGraph(csvFileName)
		} else {
			fmt.Println(errors.New("The input must be an existing csv file."))
			os.Exit(1)
		}
	}
	return models.Graph{}
}

func isValidRoute(route string) bool {
	var validRoute = regexp.MustCompile(`[a-zA-Z]{3}-[a-zA-Z]{3}`)

	return len(route) > 0 && validRoute.MatchString(route)
}
