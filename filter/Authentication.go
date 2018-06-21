package filter

import (
	"beego-api-demo/controllers"
	"fmt"

	"github.com/astaxie/beego/context"
)

var AuthFilter = func(ctx *context.Context) {
	token := ctx.Input.Header("auth")
	fmt.Println("token::::::", token)

	if token == "" {
		result := controllers.Reponse(4000, "err", "err")
		ctx.Output.JSON(result, true, true)
	}

	//models.User{}

	//if len(strToken) != 0 {
	//	myUser := models.User{Token: strToken}
	//	if myUser.Read("Token") != nil {
	//		ctx.Input.SetData("email", myUser.Email)
	//		ctx.Input.SetData("id", myUser.Id)
	//		if myUser.IsTest == true {
	//			beego.AppConfig.Set("", "")
	//			beego.AppConfig.Set("::password", "")
	//		}
	//		return
	//	}
	//}

}
