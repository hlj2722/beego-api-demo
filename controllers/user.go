package controllers

import (
	"beego-api-demo/models"
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	BaseController
}

// @Title 用户注册
// @Description 用户注册 http://localhost:8080/api/v1/user/1/register
// @Param   username
// @Param   password
// @Success 2000
// @Failure 4001 User not register
// @router / [post]
func (this *UserController) Register() {
	result := DataResponse{}
	userForm := models.UserForm{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &userForm)
	userMod := models.User{}
	userMod.Username = userForm.UserName
	userMod.Password = userForm.PassWord
	if err := userMod.Insert(); err != nil {
		result = Reponse(4000, "", "username or password error")
	} else {
		result = Reponse(2000, "", "注册成功")
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title 用户登陆
// @Description 用户登陆 http://localhost:8080/api/v1/user/1/update
// @Param   username
// @Param   password
// @Success 2000
// @Failure 4001 User not found
// @router / [post]
func (this *UserController) Login() {
	result := DataResponse{}
	userForm := models.UserForm{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &userForm)
	valid := validation.Validation{}
	b, err := valid.Valid(&userForm)
	if err != nil {
		return
	}
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	userMod := &models.User{Username: userForm.UserName, Password: userForm.PassWord}
	uId, err := userMod.Read("Username", "Password")
	if err == nil {
		tokenStr, err := this.GenToken(uId)
		if err != nil {
			beego.Debug(err.Error())
		}
		result = Reponse(2000, tokenStr, "")
	} else {
		result = Reponse(4001, "", "username or password error")
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title token验证
// @Description jwt用户验证测试，需要传入header - auth参数  http://localhost:8080/api/v1/user/auth
// @Param   header key: auth
// @Success 2000
// @Failure 4004 User not found
// @router / [get]
func (this *UserController) Auth() {
	result := DataResponse{}
	isValid, err := this.ValidToken()
	if isValid {
		result = Reponse(2000, "", "auth success")
	} else {
		result = Reponse(4001, "", err.Error())
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title 获取所有用户数据
// @Description 获取所有用户数据 http://localhost:8080/api/v1/user
// @Success 2000
// @Failure 4004 User not found
// @router / [get]
func (this *UserController) GetAll() {
	result := DataResponse{}
	userMod := models.User{}
	list := userMod.GetAllUser()
	if len(list) == 0 {
		result = Reponse(4004, "", "User not found")
	} else {
		result = Reponse(2000, list, "")
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title 更新某用户信息
// @Description 获取所有用户数据 http://localhost:8080/api/v1/user/1/update
// @Success 2000
// @Failure 4004 User not found
// @router / [post]
func (this *UserController) Update() {
	result := DataResponse{}
	uid, _ := this.GetInt64(":id")
	userMod := models.User{Id: uid}
	userMod.Phone = "138888888888"
	if err := userMod.Update("Phone"); err != nil {
		result = Reponse(4005, "", "update fail")
	} else {
		result = Reponse(2000, "", "update success")
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title 获取某个用户的信息
// @Description 获取所有用户数据 http://localhost:8080/api/v1/user/1
// @Success 2000
// @Failure 4004 User not found
// @router / [get]
func (this *UserController) GetOne() {
	result := DataResponse{}
	uid, _ := this.GetInt64(":id")
	userMod := models.User{Id: uid}
	if userInfo, err := userMod.GetUserById(uid); err != nil {
		result = Reponse(4006, "", "get user fail")
	} else {
		result = Reponse(2000, userInfo, "get user success")
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title 删除某个用户信息
// @Description 删除某个用户信息 http://localhost:8080/api/v1/user/1/del
// @Success 2000
// @Failure 4004 del user err
// @router / [post]
func (this *UserController) Delete() {
	result := DataResponse{}
	uid, _ := this.GetInt64(":id")
	userMod := models.User{Id: uid}
	if err := userMod.Delete(); err != nil {
		result = Reponse(4006, "", "del user fail")
	} else {
		result = Reponse(2000, "", "del user success")
	}
	this.Data["json"] = result
	this.ServeJSON()
}
