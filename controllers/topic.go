package controllers

import "github.com/astaxie/beego"

type TopicController struct {
	beego.Controller
}

func (t *TopicController)Get() {
	t.Data["IsTopic"]=true
	t.TplName="topic.html"
}

func (t *TopicController)Add() {
	t.TplName="topic_add.html"
	t.Ctx.WriteString("add")
}
