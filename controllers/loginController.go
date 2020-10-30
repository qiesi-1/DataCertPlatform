package controllers

import (
	"data/models"
	"github.com/astaxie/beego"
	"strings"
)

type Login struct {
	beego.Controller
}

// 直接跳转展示用户登录页面
func (l *Login) Get() {
	l.TplName = "login.html"
}

//post 方法处理用户的登录请求
func (l *Login) Post() {
	//1、解析客户端用户提交的登录数据
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("用户登录信息解析失败，请重试")
		return
	}
	// 2、根据解析到的数据，执行数据库查询操作
	u, err := user.QueryUser()
	//3、判断数据库的查询结果
	if err != nil {
		//
		l.Ctx.WriteString("用户登录失败，请重试")
		return
	}
	//3.1增加 判断用户是否实名认证，若无，跳转认证页面
	if strings.TrimSpace(u.Name) == ""||strings.TrimSpace(u.Card) =="" {//两者有其一为非实名
		l.Data["Phnoe"] = user.Phone
		l.TplName = "user_kyc.html"
		return
	}
	//4、根据查询结果返回客户端相应信息或页面跳转
	l.Data["Phone"] = u.Phone //动态数据设置
	l.TplName = "home.html"
}
