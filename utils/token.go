package utils

import (
	"beego-api-demo/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func ParaseToken(authorization string) (*models.MyCustomClaims, bool, error) {
	token, _ := jwt.ParseWithClaims(authorization, &models.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(beego.AppConfig.String("authKey")), nil
	})
	if claims, ok := token.Claims.(*models.MyCustomClaims); ok && token.Valid {
		fmt.Println("claims:", claims)
		return claims, true, nil
	}
	return nil, false, errors.New("token invalid")
}

func GenToken(uId int64) (string, error) {
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	claims := models.MyCustomClaims{
		uId,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "6617.com",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(beego.AppConfig.String("authKey")))
}
