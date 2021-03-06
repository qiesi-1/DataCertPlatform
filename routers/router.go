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
	//短信验证码登录
	beego.Router("/login_sms.html", &controllers.LoginSmsController{})
	//文件上传功能
	beego.Router("/upload", &controllers.UploadFileController{})
	//直接请求文件上传页面
	beego.Router("/home.html",&controllers.UploadFileController{})
	//查看认证数据证书页面
	beego.Router("/cert_detail.html",&controllers.CertDetailController{})
	//用户实名认证请求
	beego.Router("/user_kyc",&controllers.UserKycController{})
	//
	beego.Router("/login_sms.html",&controllers.SendSmsController{})
}
