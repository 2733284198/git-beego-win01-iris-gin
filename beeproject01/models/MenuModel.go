package models

import (
	_ "github.com/astaxie/beego/orm"
)

type MenuModel struct {
	Id     int
	Mid    int `orm: "pk:auto" `
	Parent int
	Name   string `orm :"size(45)" `
	Seq    int
	Format string `orm: "size(2048);default" `
}

type Menutree struct {
	MenuModel
	//child [] MenuModel
}

func (m *MenuModel) tableName() string {
	return "go=page"
}
