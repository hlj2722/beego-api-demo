package controllers

import (
	"time"

	"beego-api-demo/models"

	"strings"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type BaseController struct {
	beego.Controller
}

type DataResponse struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	ServerTime string      `json:"serverTime"`
}

func Reponse(errCode int, data interface{}, msg string) DataResponse {
	resp := DataResponse{
		Code:       errCode,
		Msg:        msg,
		Data:       data,
		ServerTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	return resp
}

//生成token
func (this *BaseController) GenToken(uId int64) (string, error) {
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

//验证token
func (this *BaseController) ValidToken() (bool, error) {
	authorization := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))
	if authorization == "" {
		return false, errors.New("Authorization is empty")
	}
	token, _ := jwt.ParseWithClaims(authorization, &models.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(beego.AppConfig.String("authKey")), nil
	})
	if claims, ok := token.Claims.(*models.MyCustomClaims); ok && token.Valid {
		fmt.Println("claims:", claims)
		return true, nil
	}
	return false, errors.New("Authorization invalid")
}
