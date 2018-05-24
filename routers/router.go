package routers

import (
	"easyncv/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.StartController{})
    beego.Router("/stop", &controllers.StopController{})
	beego.Router("/hcl", &controllers.HclController{})
	beego.Router("/testnomad", &controllers.TestNomadController{})
	beego.Router("/testconsul", &controllers.TestConsulController{})
}
