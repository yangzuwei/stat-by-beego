package controllers

import (
	"student/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type Foo struct {
	Name string `json:"name"`
	Num  int    `json:"num"`
}

type user struct {
	username string
}

func (this *LoginController) Get() {

	mystruct := Foo{Name: "go"}

	v := this.GetSession("asta")
	if v == nil {
		this.SetSession("asta", int(1))
		mystruct.Num = 0
	} else {
		this.SetSession("asta", v.(int)+1)
		mystruct.Num = v.(int)
	}
	this.Ctx.SetCookie("name", "magoo")
	this.Data["json"] = &mystruct
	this.ServeJSON()
}

func (this *LoginController) Post() {
	var out string
	username := this.Input().Get("username")
	password := this.GetString("password")

	if models.CheckUser(username, password) == true {
		out = "success"
	} else {
		out = "fail"
	}
	this.Ctx.WriteString(out)
}
