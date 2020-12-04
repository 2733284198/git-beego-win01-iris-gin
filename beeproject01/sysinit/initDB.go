package sysinit

import (
	"beeproject01/models"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go?charset=utf8")
	orm.RegisterModel(new(models.Page))
}
