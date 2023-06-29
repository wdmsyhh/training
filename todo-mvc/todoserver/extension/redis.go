package extension

import "github.com/gomodule/redigo/redis"

var (
	RedisClient redis.Conn
)

func InitRedis() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("123456"))
	if err != nil {
		panic(err)
	}
	RedisClient = conn
}
