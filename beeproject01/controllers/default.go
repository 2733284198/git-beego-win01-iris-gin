package controllers

import (
	"beeproject01/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	p := models.Page{
		Id:      1,
		Website: "263.com",
		Email:   "manlan@263.com",
	}

	//c.Data["Website"] = "beego.me"
	c.Data["Website"] = p.Website
	//c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Email"] = p.Email

	c.TplName = "index.tpl"
}
