package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"wwwpigcom/beeblogger/models"
	_ "wwwpigcom/beeblogger/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run("127.0.0.1:8081")
	beego.SetLevel(4)
}
