package controllers

import (
	"beeproject01/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//beeLogger "github.com/beego/bee/logger"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	/*p := models.Page{
		Id:      1,
		Website: "263.com",
		Email:   "manlan@263.com",
	}*/

	c.SetSession("username", "bingshui")
	fmt.Println(c.GetSession("username"))

	// log
	logs.Informational("====>这是log信息")

	// 先更新，在取值

	models.UpdatePage()
	p := models.GetPage()
	//c.Data["Website"] = "beego.me"
	c.Data["Website"] = p.Website
	//c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Email"] = p.Email

	c.TplName = "index.tpl"
}
