package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	//插入 10 个学生
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		id := rand.Int()
		//插入学生数据
		_, err := c.Do("HMSET", "student:"+strconv.Itoa(id), "id", id, "name", getRandomName(), "age", rand.Intn(100))
		if err != nil {
			fmt.Println("redis hmset failed:", err)
		}
		//插入投票次数
		_, err = c.Do("ZADD", "vote", rand.Intn(10), id)
		if err != nil {
			fmt.Println("redis set failed:", err)
		}
	}
	//按投票次数从低到高顺序输出
	vals,_ := redis.Values(c.Do("ZRANGE", "vote", 0, -1, "withscores"))
	i := 1
	for _, v := range vals {
		fmt.Printf("%s ", v)
		if i % 2 == 0 {
			fmt.Println()
		}
		i++
	}
}
//生成随机姓名
func  getRandomName() string {
	str := "abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}