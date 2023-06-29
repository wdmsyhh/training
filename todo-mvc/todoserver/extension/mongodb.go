package extension

import (
	"context"

	"github.com/qiniu/qmgo"
)

var (
	MongodbClient *qmgo.QmgoClient
)

func InitMongodb() {
	ctx := context.Background()
	conf := &qmgo.Config{
		Uri:      "mongodb://admin:123456@127.0.0.1:27017",
		Database: "todos",
	}
	c, err := qmgo.Open(ctx, conf)
	if err != nil {
		panic(err)
	}
	MongodbClient = c
}

func MongoCollection(collection string) *qmgo.Collection {
	return MongodbClient.Database.Collection(collection)
}
