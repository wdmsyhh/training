package extension

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	_, err := MongodbClient.Database.Collection("user").InsertOne(ctx, primitive.M{
		"name": "aaa",
		"age":  1,
	})
	if err != nil {
		panic(err)
	}
}

func TestFindUser(t *testing.T) {
	ctx := context.Background()
	query := MongodbClient.Database.Collection("user").Find(ctx, primitive.M{
		"name": "aaa",
	})
	result := primitive.M{}
	query.One(&result)
	fmt.Println(result)
}

func TestMain(m *testing.M) {
	fmt.Println("begin")
	InitMongodb()
	InitRedis()
	m.Run()
	fmt.Println("end")
}
