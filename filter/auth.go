package filter

import (
	"beego-api-demo/controllers"
	"beego-api-demo/models"
	"beego-api-demo/utils"
	"github.com/astaxie/beego/context"
)

var AuthFilter = func(ctx *context.Context) {
	token := ctx.Input.Header("Authorization")
	if token == "" {
		result := controllers.Reponse(4000, "", "require token")
		ctx.Output.JSON(result, true, true)
	}
	if claims, isValid, err := utils.ParaseToken(token); err == nil && isValid {
		var user models.User
		userMod := models.User{}
		err := userMod.Query().Filter("id", claims.UId).One(&user)
		if err != nil {
			ctx.Output.SetStatus(401)
			out := map[string]interface{}{}
			out["msg"] = "err"
			out["code"] = 4001
			ctx.Output.JSON(out, true, true)
			return
		} else {
			ctx.Input.SetData("User", user)
			return
		}
	}
	out := map[string]interface{}{}
	out["msg"] = "token invalid"
	out["code"] = 4001
	ctx.Output.JSON(out, true, true)
	return
}
