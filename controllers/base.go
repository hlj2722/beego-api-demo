package controllers

import (
	"beego-api-demo/models"
	"beego-api-demo/utils"
	"errors"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
	User *models.User
}

type DataResponse struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	ServerTime string      `json:"serverTime"`
}

func (this *BaseController) Prepare() {
	user := this.Ctx.Input.GetData("User").(models.User)
	this.User = &user
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
	return utils.GenToken(uId)
}

//验证token
func (this *BaseController) ValidToken() (int64, bool, error) {
	authorization := strings.TrimSpace(this.Ctx.Request.Header.Get("Authorization"))
	if authorization == "" {
		return 0, false, errors.New("Authorization is empty")
	}
	if claims, isValid, err := utils.ParaseToken(authorization); err == nil && isValid {
		return claims.UId, true, nil
	}
	return 0, false, errors.New("Authorization invalid")
}
