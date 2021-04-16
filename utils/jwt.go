package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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

// func Verify(ctx *gin.Context) (*jwt.Token, error) {
// 	accessToken := ctx.GetHeader("Authorization")

// 	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
// 		return accessToken, nil
// 	})

// 	if err != nil {
// 		defer fmt.Sprintf("verified accessToken failed %v", accessToken)
// 		logrus.Fatal(err.Error())
// 		return token, err
// 	}

// 	defer fmt.Sprintf("verified accessToken success %v", accessToken)
// 	return token, err
// }
