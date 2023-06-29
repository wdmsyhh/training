package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"todoserver/extension"
	"todoserver/model"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
}

// 注册
func (userService *UserService) Register(ctx context.Context, user model.User) error {
	user.Id = primitive.NewObjectID().Hex()
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	_, err := extension.MongoCollection("user").InsertOne(ctx, user)
	return err
}

// 登录
func (userService *UserService) Login(ctx context.Context, user model.User) (*model.User, error) {
	password := user.Password
	err := extension.MongoCollection("user").Find(ctx, bson.M{"userName": user.UserName}).One(&user)
	if err != nil {
		return nil, err
	}
	if user.Password == fmt.Sprintf("%x", md5.Sum([]byte(password))) {
		// 用户信息存 redis
		// _, err := extension.RedisClient.Do("HMSET", fmt.Sprintf("%x", md5.Sum([]byte(user.Id))),
		// 	"id", user.Id,
		// 	"userName", user.UserName,
		// 	"password", user.Password,
		// )
		// extension.RedisClient.Do("EXPIRE", "user:"+user.Id, 30*60)
		// if err != nil {
		// 	return nil, err
		// }
		return &user, nil
	}
	return nil, nil
}
