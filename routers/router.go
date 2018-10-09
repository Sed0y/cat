package routers

import (
	"cat/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/admin", &controllers.AdminController{}, "get:Dashboard")
	beego.Router("/admin/categories", &controllers.AdminController{}, "get:Categories")

	beego.Router("/", &controllers.MainController{})
}
