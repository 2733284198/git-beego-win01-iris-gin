package controllers

import (
	"beedemo01/models"
	_ "beedemo01/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	m := models.GetPage()
	c.Data["Website"] = m.Website
	c.Data["Email"] = m.Email

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
