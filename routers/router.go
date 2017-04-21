package routers

import (
	"MockServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/MockServer/appserver", &controllers.MainController{})
}
