package model

import "github.com/globalsign/mgo/bson"

type Plan struct {
	Id          bson.ObjectId `json:"_id" bson:"_id"`
	Content     string        `json:"content" bson:"content"`
	IsCompleted bool          `json:"isCompleted" bson:"isCompleted"`
	UserId      string        `json:"userId" bson:"userId"`
}
