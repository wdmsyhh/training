package extension

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestGenerateAccessToken(t *testing.T) {
	s, err := GenerateAccessToken("123", "aaa", time.Second*15)
	if err != nil {
		panic(err)
	}
	log.Println(s)
}

func TestParseAccessToken(t *testing.T) {
	tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhaWQiOiIiLCJhdWQiOiIiLCJleHAiOjE2ODgwMDgzMzIsImlhdCI6MTY4ODAwODMxNywiaXNzIjoidG9kb1NlcnZlciIsInN1YiI6ImFhYSJ9.VwufERweuKEbYTzErACCshAyy-tlKAyH0VH_LGR_gQE"
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		// TODO 一般是存到数据库或 redis 中的，从数据库或 redis 获取
		return []byte("123"), nil
	}
	token, err := ParseAccessToken(tokenStr, keyFunc)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
