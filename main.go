package main

import (
	_ "student/routers"
	"student/views"

	"github.com/astaxie/beego"
)

func main() {

	//beego.BConfig.WebConfig.Session.SessionOn = true
	beego.AddFuncMap("add", views.Add)
	beego.AddFuncMap("div", views.Div)
	beego.AddFuncMap("mul", views.Mul)
	beego.Run()
}
