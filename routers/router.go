package routers

import (
	"data/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/register", &controllers.RegisterController{})
}

