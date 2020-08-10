package entities

import "gopkg.in/mgo.v2/bson"

type Url struct {
	Id  bson.ObjectId `bson:"_id"`
	Url string        `bson:"url"`
}
