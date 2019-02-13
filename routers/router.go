package routers

import (
	"context"
	"student/controllers"

	"github.com/astaxie/beego"
)

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
}
