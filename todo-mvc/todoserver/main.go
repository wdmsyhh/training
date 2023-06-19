package main

import (
	"fmt"
	"log"
	"todoserver/controller"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cast"
)

func main() {
	// router := gin.Default()
	// registerRouter(router)
	// router.Run(":8088")

	c, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("123456"))
	if err != nil {
		panic(err)
	}
	reply, err2 := c.Do("GET", "a")
	if err != nil {
		log.Println(err2.Error())
	}
	fmt.Println(cast.ToString(reply))
}

// 路由设置
func registerRouter(router *gin.Engine) {
	new(controller.PlansController).Router(router)
	new(controller.UserController).Router(router)
}
