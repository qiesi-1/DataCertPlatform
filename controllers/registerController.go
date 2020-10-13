package controllers

import (
	"data/models"
	"github.com/beego"
)

type RegisterController struct {
	beego.Controller
}

//处理用户注册
func (r *RegisterController) Post(){
	var user models.User
	err := r.ParseForm(&user)
	if err != nil{
		r.Ctx.WriteString("抱歉，数据解析失败，请重试")
		return
	}
	// 2、将解析到的数据保存到数据库中
	_ ,err = user.AddUser()
	if err!= nil {
		r.Ctx.WriteString("抱歉，用户注册失败，请重试")
		return
	}
	//3.将处理结果返回给浏览器客户端
	//3.1 成功，跳转登录页面
	r.TplName = "login.html"
}