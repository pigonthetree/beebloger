package routers

import (
	"github.com/astaxie/beego"
	"wwwpigcom/beeblogger/controllers"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.AutoRouter(&controllers.TopicController{})
    beego.Router("/topic", &controllers.TopicController{})
}
