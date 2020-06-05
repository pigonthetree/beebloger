package controllers

import (
	"github.com/astaxie/beego"
	"wwwpigcom/beeblogger/models"
)

type TopicController struct {
	beego.Controller
}

func (t *TopicController) Get() {
	t.Data["IsLogin"] = checkAccount(t.Ctx)
	t.Data["IsTopic"]=true
	t.TplName="topic.html"
}

func (t *TopicController) Post()  {
	if !checkAccount(t.Ctx){
		t.Redirect("/login",302)
		return
	}
	title:=t.Input().Get("title")
	content:=t.Input().Get("content")
	var err error
	err = models.AddTopic(title, content)
	if err!=nil{
		beego.Error(err)
	}
	t.Redirect("/topic",302)
}

func (t *TopicController)Add() {
	t.TplName="topic_add.html"
}
