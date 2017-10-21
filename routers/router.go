package routers

import (
	"easyncv/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.StartController{})
    beego.Router("/stop", &controllers.StopController{})
}
