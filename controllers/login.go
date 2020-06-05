package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	isExit:=l.Input().Get("exit")=="true"
	if isExit {
		l.Ctx.SetCookie("uname","",-1,"/")
		l.Ctx.SetCookie("pwd","",-1,"/")
		l.Redirect("/",302)
		return
	}
	l.TplName = "login.html"
}

func (l *LoginController) Post() {
	uname:=l.Input().Get("uname")
	pwd:=l.Input().Get("pwd")
	autoLogin:=l.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname")==uname && beego.AppConfig.String("pwd")==pwd {
		maxAge:=0
		if autoLogin{
			maxAge = 1<<31 -1
		}
		// 设置cookie
		l.Ctx.SetCookie("uname",uname,maxAge,"/")
		l.Ctx.SetCookie("pwd",pwd,maxAge,"/")
	}

	l.Redirect("/",302)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil{
		return false
	}
	uname:=ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err!=nil{
		return false
	}
	pwd := ck.Value
	return beego.AppConfig.String("uname")==uname && beego.AppConfig.String("pwd")==pwd
}
