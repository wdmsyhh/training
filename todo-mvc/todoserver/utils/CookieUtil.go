package utils

import (
	"github.com/gin-gonic/gin"
)

type CookieUtil struct {
}

func (cookieUtil *CookieUtil) GetCookie(context *gin.Context) string {
	userKey, err := context.Cookie("userKey")
	if err != nil {
		return ""
	}
	return userKey
}
