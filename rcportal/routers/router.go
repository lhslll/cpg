package routers

import (
	"github.com/cpg/rcportal/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
}
