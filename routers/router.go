// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego-api-demo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/user",
			beego.NSRouter("/", &controllers.UserController{}, "get:GetAll"),
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
			beego.NSRouter("/login/", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/auth/", &controllers.UserController{}, "get:Auth"),
			beego.NSRouter("/:id:int/update", &controllers.UserController{}, "post:Update"),
			beego.NSRouter("/:id:int/", &controllers.UserController{}, "get:GetOne"),
			beego.NSRouter("/:id:int/del", &controllers.UserController{}, "post:Delete"),
		),
	)
	beego.AddNamespace(ns)
}
