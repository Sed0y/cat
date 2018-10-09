package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Dashboard() {

	c.Layout = "layout/default.tpl"
	c.TplName = "admin/statistics.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Navigation"] = "element/navigation.tpl"

}

func (c *AdminController) Categories() {

	c.Data["Categories"] = App.Categories.List

	c.Layout = "layout/default.tpl"
	c.TplName = "category/categories.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Navigation"] = "element/navigation.tpl"

}
