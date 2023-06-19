package utils

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedisUtil struct {
}

func (redisUtil *RedisUtil) ConnectRedis() redis.Conn {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil
	}
	return conn
}

func (redisUtil *RedisUtil) GetUserIdFromRedis(userKey string) string {
	conn := redisUtil.ConnectRedis()
	defer conn.Close()

	userId, _ := conn.Do("hget", userKey, "id")

	bt := []byte{}
	if userId != nil {
		for _, b := range userId.([]uint8) {
			bt = append(bt, byte(b))
		}
	}
	return string(bt)
}
