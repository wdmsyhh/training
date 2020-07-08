package utils

import (
	"fmt"

	"github.com/globalsign/mgo"
)

type MongoUtil struct {
}

//连接 mongo
func (mongoUtil *MongoUtil) ConnectMongo(collectionName string) (*mgo.Session, *mgo.Collection) {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect success")
	}
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("todos").C(collectionName)
	return session, collection
}
