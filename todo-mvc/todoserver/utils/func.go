package utils

import (
	"todoserver/extension"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromRedis(userKey string) string {
	userId, _ := extension.RedisClient.Do("hget", userKey, "id")
	bt := []byte{}
	if userId != nil {
		for _, b := range userId.([]uint8) {
			bt = append(bt, byte(b))
		}
	}
	return string(bt)
}

func GetUserKey(context *gin.Context) string {
	userKey, err := context.Cookie("userKey")
	if err != nil {
		return ""
	}
	return userKey
}
