package router

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	. "github.com/yasmindias/travelhelper/models"
	"github.com/yasmindias/travelhelper/utils"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	file := utils.OpenFileToRead()
	routes := utils.ReadFile(file)
	respondWithJson(w, http.StatusOK, routes)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var route Route
	if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	file := utils.OpenFileToWrite()
	err := utils.WriteToFile(file, route)

	if err == nil {
		respondWithJson(w, http.StatusCreated, route)
	}
}

func FindBestRoute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	graph := utils.StartGraphWithCsvFile(os.Getenv("filename"))
	path := graph.Dijkstra(params["origin"], params["destiny"])

	err := json.NewEncoder(w).Encode(&path)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, "Couldn't find best route")
		return
	}
}
