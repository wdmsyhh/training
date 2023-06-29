package extension

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateAccessToken(secretKey, userId string, tokenExpireTime time.Duration) (string, error) {
	method := jwt.GetSigningMethod("HS256")
	token := jwt.New(method)

	iat := time.Now().Unix()
	exp := iat + int64(tokenExpireTime.Seconds())

	claims := jwt.MapClaims{
		"iat": iat,
		"exp": exp,
		"iss": "todoServer",
		"aud": "",
		"sub": userId,
		"aid": "",
	}

	header := map[string]interface{}{
		"typ": "JWT",
		"alg": method.Alg(),
	}

	token.Claims = claims
	token.Header = header
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ParseAccessToken(accessToken string, lookup jwt.Keyfunc) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, lookup)
	if err != nil {
		return nil, err
	}
	return token, nil
}
