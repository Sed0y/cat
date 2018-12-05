package routers

import (
	"cat/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/admin", &controllers.AdminController{}, "get:Dashboard")
	beego.Router("/admin/categories", &controllers.AdminController{}, "get:Categories")
	beego.Router("/admin/categories/add", &controllers.AdminController{}, "post:AddCategory")

	beego.Router("/", &controllers.MainController{})
}
