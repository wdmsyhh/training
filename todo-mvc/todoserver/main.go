package main

import (
	"todoserver/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	registerRouter(router)
	router.Run(":8088")
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(controller.PlansController).Router(router)
	new(controller.UserController).Router(router)
}
