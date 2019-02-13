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
	this.TplName = "login.html"
}

func (this *LoginController) Post() {

	username := this.Input().Get("username")
	password := this.GetString("password")
	user := models.GeUserByUsername(username)
	if models.CheckUser(user, password) {

		v := this.GetSession("uid")
		if v == nil {
			this.SetSession("uid", user.Id)
		}
		this.Ctx.SetCookie("username", username)

		this.Redirect("/", 302)
	} else {
		this.Redirect("/login", 302)
	}

}
