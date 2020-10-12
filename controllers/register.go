package controllers

import (
	"data/models"
	"github.com/beego"
)

type RegisterController struct {
	beego.Controller
}


func (r *RegisterController) Post(){


	var user models.User
	err := r.ParseForm(&user)
	if err != nil{
		r.Ctx.WriteString("抱歉数据错误")
		return
	}

	r.TplName = "login.html"
}