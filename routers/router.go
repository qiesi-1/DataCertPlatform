package routers

import (
	"data/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//用户注册接口
	beego.Router("/register", &controllers.RegisterController{})
	//用户登录接口
	beego.Router("/login", &controllers.Login{})
	//请求直接登录的页面
	beego.Router("/login.html", &controllers.Login{})
	//文件上传功能
	beego.Router("/upload", &controllers.UploadFileController{})
}
