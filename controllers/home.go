package controllers

import (
	"github.com/astaxie/beego"
	"wwwpigcom/beeblogger/models"
)

type HomeController struct {
	beego.Controller
}

func (h *HomeController) Get() {
	h.Data["IsHome"] = true
	h.TplName = "home.html"

	h.Data["IsLogin"] = checkAccount(h.Ctx)
	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	} else {
		h.Data["Topics"] = topics
	}
}
