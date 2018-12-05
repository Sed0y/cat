package controllers

import (
	"cat/models/entities"
	"fmt"
	"html/template"
	"strconv"

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

	c.Data["CategoriesRender"] = template.HTML(App.Categories.RenderToAdminPanel())
	c.Data["CategoriesSelectRender"] = template.HTML(App.Categories.RenderSelectList())
	c.Data["CategoriesSelectWeight"] = template.HTML(App.Categories.RenderWeightSelectList())

	c.Layout = "layout/default.tpl"
	c.TplName = "category/categories.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Navigation"] = "element/navigation.tpl"

}

func (c *AdminController) AddCategory() {

	var current models.Category

	current.Name = c.Ctx.Request.FormValue("name")
	current.ParentId, _ = strconv.Atoi(c.Ctx.Request.FormValue("parent_id"))
	current.Weight, _ = strconv.Atoi(c.Ctx.Request.FormValue("wheight"))
	current.Active, _ = strconv.ParseBool(c.Ctx.Request.FormValue("active"))
	current.Url = c.Ctx.Request.FormValue("url")

	fmt.Println(current)

	res := current.Create()

	if res == true {
		App.Categories.Load()
		c.Ctx.ResponseWriter.Write([]byte("ok"))
	} else {
		c.Ctx.ResponseWriter.Write([]byte("err"))
	}

	/*
		c.Data["CategoriesRender"] = template.HTML(App.Categories.RenderToAdminPanel())
		c.Data["CategoriesSelectRender"] = template.HTML(App.Categories.RenderSelectList())
		c.Data["CategoriesSelectWeight"] = template.HTML(App.Categories.RenderWeightSelectList())

		c.Layout = "layout/default.tpl"
		c.TplName = "category/categories.tpl"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Navigation"] = "element/navigation.tpl"
	*/

}
