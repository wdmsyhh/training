package controller

import (
	"net/http"
	"time"
	"todoserver/extension"
	"todoserver/model"
	"todoserver/response"
	"todoserver/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	returnUser, err := userService.Login(context, user)
	if returnUser == nil || err != nil {
		response.Failed(context, response.LOGIN_FAILED)
		return
	}
	//设置 cookie
	account := make(primitive.M)
	err = extension.MongoCollection("account").Find(context, primitive.M{"enable": true}).One(account)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "find account err:" + err.Error(),
		})
		return
	}
	expireTime := time.Duration(time.Second * 30)
	accessToken, err := extension.GenerateAccessToken(account["secretKey"].(string), returnUser.Id, expireTime)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}
	context.SetCookie("accessToken", accessToken, int(expireTime.Seconds()), "/", "localhost", false, true)

	response.Success(context, returnUser)
}

func (userController *UserController) register(context *gin.Context) {
	var user model.User
	err := context.BindJSON(&user)
	if err != nil {
		response.Failed(context, response.BIND_PARAM_ERROR)
		return
	}
	err = userService.Register(context, user)
	if err != nil {
		response.Failed(context, response.REGISTER_FAILED)
		return
	}
	response.Success(context, response.REGIDTER_SUCCESS)
}
