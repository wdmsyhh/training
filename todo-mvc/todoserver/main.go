package main

import (
	"todoserver/controller"
	"todoserver/extension"
	"todoserver/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	extension.InitMongodb()
	extension.InitRedis()
	router := gin.Default()
	registerRouter(router)
	router.Run(":8088")
}

// 路由设置
func registerRouter(router *gin.Engine) {
	authGroup := router.Group("")
	authGroup.Use(middleware.Auth())
	new(controller.PlansController).Router(authGroup)
	new(controller.UserController).Router(router)
}
