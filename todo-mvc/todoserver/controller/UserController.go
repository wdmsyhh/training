package controller

import (
	"crypto/md5"
	"fmt"
	"todoserver/model"
	"todoserver/response"
	"todoserver/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

var userService = new(service.UserService)

func (userController *UserController) Router(router *gin.Engine) {
	//登录
	router.POST("/user/login", userController.login)
	//注册
	router.POST("/user/register", userController.register)
}

func (userController *UserController) login(context *gin.Context) {
	var user model.User
	err := context.BindJSON(&user)
	if err != nil {
		response.Failed(context, response.LOGIN_FAILED)
		return
	}
	returnUser, err := userService.Login(user)
	if returnUser == nil || err != nil {
		response.Failed(context, response.LOGIN_FAILED)
		return
	}
	//设置 cookie
	context.SetCookie("userKey", fmt.Sprintf("%x", md5.Sum([]byte(returnUser.Id))), 30*60, "/", "localhost", false, true)

	response.Success(context, returnUser)
}

func (userController *UserController) register(context *gin.Context) {
	var user model.User
	err := context.BindJSON(&user)
	if err != nil {
		response.Failed(context, response.BIND_PARAM_ERROR)
		return
	}
	err = userService.Register(user)
	if err != nil {
		response.Failed(context, response.REGISTER_FAILED)
		return
	}
	response.Success(context, response.REGIDTER_SUCCESS)
}
