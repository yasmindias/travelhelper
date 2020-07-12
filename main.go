package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
	router "github.com/yasmindias/travelhelper/router"
	"github.com/yasmindias/travelhelper/utils"
)

func main() {
	if len(os.Args) > 1 {
		runCommandLine(os.Args[1])
	} else {
		runHttpServer()
	}
}

func runCommandLine(filename string) {
	graph := utils.StartGraphWithCsvFile(filename)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter the route: ")
	route, _ := reader.ReadString('\n')
	if isValidRoute(route) {
		places := strings.Split(strings.ToUpper(route), "-")
		origin := strings.Trim(places[0], " ")
		destiny := strings.TrimSuffix(places[1], "\n")

		path := graph.Dijkstra(origin, destiny)
		fmt.Println("Best route: " + graph.PrintPath(path))
	} else {
		fmt.Println(errors.New("The input must be in the format \"ORG-DEST\"."))
	}
}

func runHttpServer() {
	r := mux.NewRouter()
	r.HandleFunc("api/routes", router.GetAll).Methods("GET")
	r.HandleFunc("api/routes", router.GetAll).Methods("POST")
	r.HandleFunc("api/bestroute", router.GetAll).Methods("GET")

	port := ":3000"
	fmt.Println("Server running in port: ", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func isValidRoute(route string) bool {
	var validRoute = regexp.MustCompile(`[a-zA-Z]{3}-[a-zA-Z]{3}`)

	return len(route) > 0 && validRoute.MatchString(route)
}
