package router

import (
	"encoding/json"
	"net/http"

	. "github.com/yasmindias/travelhelper/config/dao"
	. "github.com/yasmindias/travelhelper/models"
	"gopkg.in/mgo.v2/bson"
)

var dao = RouterDAO{}

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
	routes, err := dao.GetAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, routes)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var route Route
	if err := json.NewDecoder(r.Body).Decode(&route); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	route.ID = bson.NewObjectId()
	if err := dao.Create(route); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, route)
}
