package dao

import (
	"log"

	"github.com/yasmindias/travelhelper/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RouterDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "routes"
)

func (r *RouterDAO) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(r.Database)
}

func (e *RouterDAO) GetAll() ([]models.Route, error) {
	var routes []models.Route
	err := db.C(COLLECTION).Find(bson.M{}).All(&routes)
	return routes, err
}

func (e *RouterDAO) Create(route models.Route) error {
	err := db.C(COLLECTION).Insert(&route)
	return err
}
