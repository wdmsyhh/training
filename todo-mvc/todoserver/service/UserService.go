package service

import (
	"crypto/md5"
	"fmt"
	"todoserver/model"
	"todoserver/utils"

	"github.com/globalsign/mgo/bson"
)

type UserService struct {
}

var redisUtil = new(utils.RedisUtil)

//注册
func (userService *UserService) Register(user model.User) error {
	session, c := mongoUtil.ConnectMongo("user")
	defer session.Close()

	user.Id = bson.NewObjectId().Hex()
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	err := c.Insert(user)
	return err
}

//登录
func (userService *UserService) Login(user model.User) (*model.User, error) {
	session, c := mongoUtil.ConnectMongo("user")
	defer session.Close()

	password := user.Password
	err := c.Find(bson.M{"userName": user.UserName}).One(&user)
	if err != nil {
		return nil, err
	}
	if user.Password == fmt.Sprintf("%x", md5.Sum([]byte(password))) {
		//用户信息存 redis
		conn := redisUtil.ConnectRedis()
		defer conn.Close()

		_, err := conn.Do("HMSET", fmt.Sprintf("%x", md5.Sum([]byte(user.Id))), "id", user.Id, "userName", user.UserName, "password", user.Password)
		conn.Do("EXPIRE", "user:"+user.Id, 30*60)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, nil
}
