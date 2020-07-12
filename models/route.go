package models

import "gopkg.in/mgo.v2/bson"

type Route struct {
	ID      bson.ObjectId `bson: "_id" json: id`
	origin  string        `bson: "origin" json: "origin`
	destiny string        `bson: "destiny" json: "destiny`
}
