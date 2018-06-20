package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

type baseController struct {
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
