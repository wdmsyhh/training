package middleware

import (
	"net/http"
	"todoserver/extension"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("accessToken")
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "get token from cookie err:" + err.Error(),
			})
			c.Abort()
			return
		}
		if accessToken == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "no access token",
			})
			c.Abort()
			return
		}

		keyFunc := func(t *jwt.Token) (interface{}, error) {
			var secretKey string
			reply, err := extension.RedisClient.Do("GET", "jwt_secretKey")
			if err != nil {
				secretKey = cast.ToString(reply)
			}
			if secretKey == "" {
				account := make(primitive.M)
				err = extension.MongoCollection("account").Find(c, primitive.M{"enable": true}).One(account)
				if err != nil {
					return nil, err
				}
				secretKey = account["secretKey"].(string)
				extension.RedisClient.Do("SET", "jwt_secretKey", secretKey)
			}
			return []byte(secretKey), nil
		}

		token, err := extension.ParseAccessToken(accessToken, keyFunc)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    0,
				"message": "parse token err:" + err.Error(),
				"token":   token,
			})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["sub"]
		c.Set("userId", userId)
		c.Next()
	}
}
