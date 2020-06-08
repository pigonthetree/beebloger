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
	t.Data["IsTopic"] = true
	t.TplName = "topic.html"
	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	} else {
		t.Data["Topic"] = topics
	}
}

func (t *TopicController) Post() {
	if !checkAccount(t.Ctx) {
		t.Redirect("/login", 302)
		return
	}
	tid := t.Input().Get("tid")
	title := t.Input().Get("title")
	content := t.Input().Get("content")
	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, content)
	} else {
		err = models.ModifyTopic(tid, title, content)
	}
	if err != nil {
		beego.Error(err)
	}
	t.Redirect("/topic", 302)
}

func (t *TopicController) Add() {
	t.TplName = "topic_add.html"
}

func (t *TopicController) View() {
	t.TplName = "topic_view.html"
	topic, err := models.GetTopic(t.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		t.Redirect("/", 302)
		return
	}
	t.Data["Topic"] = topic
	t.Data["Tid"] = t.Ctx.Input.Param("0")
}

func (t *TopicController) Modify() {
	t.TplName = "topic_modify.html"
	tid := t.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		t.Redirect("/", 302)
	}
	t.Data["Topic"] = topic
	t.Data["Tid"] = tid
}
