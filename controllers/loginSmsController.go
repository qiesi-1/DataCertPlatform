package controllers

import "github.com/astaxie/beego"

type LoginSmsController struct {
	beego.Controller
}

func(l *LoginSmsController)Get(){
	l.TplName = "login_sms.html"
}

//短信验证码登录
func(l *LoginSmsController)Post(){

}
