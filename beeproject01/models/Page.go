package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	Id      int
	Website string
	Email   string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go?charset=utf8")
	orm.RegisterModel(new(Page))
}

func GetPage() Page {
	//trn := Page{Website: "baidu.com", Email: "manlan2010@163.com"}
	//return trn

	o := orm.NewOrm()
	p := Page{Id: 1}

	err := o.Read(&p)
	fmt.Println(&err)

	return p
}

func UpdatePage() {
	p := Page{
		Id:      1,
		Website: "jjdd.com",
		Email:   "jjdd@jd.com",
	}

	o := orm.NewOrm()
	o.Update(&p)
}
