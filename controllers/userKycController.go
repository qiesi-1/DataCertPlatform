package controllers

import (
	"data/models"
	"github.com/astaxie/beego"
)

type UserKycController struct {
	beego.Controller
}
//浏览器get请求，用于跳转实名认证页面
func (u *UserKycController)Get()  {
	u.TplName = "user_kyc.html"
}


//form表单的post请求
func (u *UserKycController)Post()  {
	var user models.User
	err := u.ParseForm(&user)
	if err!=nil {
		u.Ctx.WriteString("baoqian,数据解析错误")
		return
	}

	//2.把实名认证更新到数据库用户表中
	_,err = user.UpdateUser()
	//3.判断实名认证结果
	if err!=nil {
		u.Ctx.WriteString("30抱歉，用户实名认证失败，请重试")
		return
	}
	//4.跳转
	u.TplName = "home.html"
}