package utils

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	UId int64 `json:"uid"`
	jwt.StandardClaims
}

func ParaseToken(authorization string) (*MyCustomClaims, bool, error) {
	token, _ := jwt.ParseWithClaims(authorization, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(beego.AppConfig.String("authKey")), nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Println("claims:", claims)
		return claims, true, nil
	}
	return nil, false, errors.New("token invalid")
}

func GenToken(uId int64) (string, error) {
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	claims := MyCustomClaims{
		uId,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "6617.com",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(beego.AppConfig.String("authKey")))
}
