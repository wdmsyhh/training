package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Plan struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	Content     string             `json:"content" bson:"content"`
	IsCompleted bool               `json:"isCompleted" bson:"isCompleted"`
	UserId      string             `json:"userId" bson:"userId"`
}
