package main

import (
	_ "student/routers"
	"student/views"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	//beego.BConfig.WebConfig.Session.SessionOn = true
	beego.AddFuncMap("add", views.Add)
	beego.AddFuncMap("div", views.Div)
	beego.AddFuncMap("mul", views.Mul)
	beego.Run()
}
