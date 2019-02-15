package controllers

import (
	"fmt"
	"student/models"
	"time"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.zaisb.com"
	c.Data["Email"] = "yangzuwei@outlook.com"
	models.Stat() //调用统计文件
	c.Data["TotalNum"] = models.TotalNum()
	c.Data["HasHandle"] = models.HasHanle()

	c.Data["Accounts"] = models.Accounts
	c.Data["Actually"] = models.Actually
	c.Data["Current"] = time.Now().Format("2006-01-02")

	c.TplName = "index.tpl"
}
