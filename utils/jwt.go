package utils

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Sign(UserId uint, Email, SecrePublicKey string, ExpiredAt time.Duration) (string, error) {

	expiredAt := time.Now().Add(time.Minute * ExpiredAt).Unix()
	jwtSecretKey := SecrePublicKey

	claims := jwt.MapClaims{}
	claims["id"] = UserId
	claims["email"] = Email
	claims["expiredAt"] = expiredAt
	claims["authorization"] = true

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := to.SignedString([]byte(jwtSecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func Verify(ctx *gin.Context, SecrePublicKey string) (*jwt.Token, error) {
	tokenHeader := ctx.GetHeader("Authorization")
	accessToken := strings.SplitAfter(tokenHeader, "Bearer")[1]

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(SecrePublicKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return token, err
	}

	return token, nil
}
