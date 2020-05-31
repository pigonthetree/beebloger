package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	c.Data["TrueCondition"] = true

	type Users struct {
		Name string
		Age int
		Gender string
	}
	
	user:=Users{
		Name:   "zhangsan",
		Age:    21,
		Gender: "Male",
	}
	c.Data["Users"]=user

	c.Data["HtmlVar"]="hey gays!"
	c.Data["Html"]="<div>This is html!</div>"
	c.Data["Pipe"]="<div>This is html!</div>"
}
