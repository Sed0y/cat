package app

import (
	"cat/conf"
	"cat/models/entities"
	//"fmt"
	//"log"
	"sort"
	"strconv"

	_ "github.com/lib/pq"
)

type Categories struct {
	List []*models.Category
}

// Загружает категории из БД
func (c *Categories) Load() bool {

	//qs.OrderBy("-profile__age", "profile")
	// ORDER BY profile.age DESC, profile_id ASC
	qs := conf.AppOrm.QueryTable("category")
	qs.OrderBy("id").All(&c.List)
	//qs.All(&c.List)
	//conf.AppOrm.QueryTable("category").All(&c.List)

	return true

}

// Добавить категорию в список
// добавляет в БД, после чего обновляет список
// *
func (c *Categories) Add(
	ParentId int,
	Level int,
	Weight int,
	Name string,
	Active bool,
	Url string) bool {

	var NewOne models.Category

	NewOne.ParentId = ParentId
	NewOne.Level = Level
	NewOne.Weight = Weight
	NewOne.Name = Name
	NewOne.Active = Active
	NewOne.Url = Url

	/*
		Проверки данных ???

	*/

	result := NewOne.Create()

	if result == true {
		result = c.Load()
		if result == true {
			return true
		} else {
			return false
		}
		return true

	} else {
		return false
	}

	return true
}

// Удалить элемент из списка
// ### Внимание! Надо добавить проверку, что если этот элемент
// ### является родителем других, то удалять его нельзя
// *
func (c *Categories) Remove(CategoryID int) bool {

	// !!! Проверку на наличие детей !!!
	for i := 0; i < len(c.List); i++ {
		if c.List[i].Id == CategoryID {
			if c.List[i].Delete() {
				c.Load()
				return true
			} else {
				return false
			}
		}
	}

	return false
}

// Генерирует html со списком категорий для панели управления
// *
func (c *Categories) RenderToAdminPanel() string {

	var cats string

	cats = ""
	new_line := "\r\n"

	// Сортировка по "весу"
	sort.Slice(c.List, func(i, j int) bool {
		if c.List[i].Weight < c.List[j].Weight {
			return true
		}
		if c.List[i].Weight > c.List[j].Weight {
			return false
		}
		return true
	})

	cats += "<div id=\"categories-view-admin_panel\">" + new_line

	index_level_1 := 1
	index_level_2 := 1
	index_level_3 := 1

	for i := 0; i < len(c.List); i++ {
		if c.List[i].Level != 1 {
			continue
		}

		cats += "<div "
		cats += " class=\"category-level-1\" >"

		cats += "<span>"
		cats += strconv.Itoa(index_level_1) + "."
		cats += "</span>"

		cats += "	<a "
		cats += " data-toggle=\"collapse\" "
		cats += " href=\"#collapse_cats_" + strconv.Itoa(c.List[i].Id) + "\" "
		cats += " aria-expanded=\"false\" "
		cats += " aria-controls=\"collapse_cats_" + strconv.Itoa(c.List[i].Id) + "\" "
		cats += " id=\"catid-" + strconv.Itoa(c.List[i].Id) + "\" "
		cats += ">" + new_line

		cats += c.List[i].Name

		cats += "	</a> " + new_line

		cats += "<span class=\"cat-weight\">"
		cats += strconv.Itoa(c.List[i].Weight)
		cats += "</span>"

		cats += "<span class=\"cat-url\"><i>"
		cats += c.List[i].Url
		cats += "</i></span>"

		cats += "<span class=\"cat-action\">"
		cats += "Редактировать"
		cats += "</span>"

		cats += "	</div> " + new_line

		cats += "	<div "

		if c.ChildCountOf(c.List[i].Id) == 0 {
			cats += " class=\"category-level-1-block collapse\" "
		} else {
			cats += " class=\"category-level-1-block collapse show\" "
		}

		cats += " id=\"collapse_cats_" + strconv.Itoa(c.List[i].Id) + "\" "
		cats += "> " + new_line

		index_level_2 = 1
		for j := 0; j < len(c.List); j++ {
			if c.List[j].Level != 2 {
				continue
			} else if c.List[j].ParentId != c.List[i].Id {
				continue
			}

			cats += "<div "
			cats += " class=\"category-level-2\" >"

			cats += "<span>"
			cats += strconv.Itoa(index_level_1) + "." + strconv.Itoa(index_level_2) + ". "
			cats += "</span>"

			cats += "	<a "
			cats += " data-toggle=\"collapse\" "
			cats += " href=\"#collapse_cats_" + strconv.Itoa(c.List[j].Id) + "\" "
			cats += " aria-expanded=\"false\" "
			cats += " aria-controls=\"collapse_cats_" + strconv.Itoa(c.List[j].Id) + "\" "
			cats += " id=\"catid-" + strconv.Itoa(c.List[j].Id) + "\" "
			cats += ">" + new_line

			cats += c.List[j].Name

			cats += "	</a> " + new_line

			cats += "<span class=\"cat-weight\">"
			cats += strconv.Itoa(c.List[j].Weight)
			cats += "</span>"

			cats += "<span class=\"cat-url\"><i>"
			cats += c.List[j].Url
			cats += "</i></span>"

			cats += "<span class=\"cat-action\">"
			cats += "Редактировать"
			cats += "</span>"

			cats += "	</div> " + new_line

			if c.ChildCountOf(c.List[j].Id) != 0 {

				cats += "	<div "
				cats += " class=\"category-level-2-block collapse show\" "
				cats += " id=\"collapse_cats_" + strconv.Itoa(c.List[j].Id) + "\" "
				cats += "> " + new_line

				for k := 0; k < len(c.List); k++ {
					if c.List[k].Level != 3 {
						continue
					} else if c.List[k].ParentId != c.List[j].Id {
						continue
					}

					cats += "<div "
					cats += " class=\"category-level-3\" >"

					cats += "<span>"
					cats += strconv.Itoa(index_level_1) + "." + strconv.Itoa(index_level_2) + "." + strconv.Itoa(index_level_3) + ". "
					cats += "</span>"

					cats += "	<a "
					cats += " data-toggle=\"collapse\" "
					cats += " href=\"#collapse_cats_" + strconv.Itoa(c.List[k].Id) + "\" "
					cats += " aria-expanded=\"false\" "
					cats += " aria-controls=\"collapse_cats_" + strconv.Itoa(c.List[k].Id) + "\" "
					cats += " id=\"catid-" + strconv.Itoa(c.List[k].Id) + "\" "
					cats += ">" + new_line

					cats += c.List[k].Name

					cats += "	</a> " + new_line

					cats += "<span class=\"cat-weight\">"
					cats += strconv.Itoa(c.List[k].Weight)
					cats += "</span>"

					cats += "<span class=\"cat-url\"><i>"
					cats += c.List[k].Url
					cats += "</i></span>"

					cats += "<span class=\"cat-action\">"
					cats += "Редактировать"
					cats += "</span>"

					cats += "</div>"

					index_level_3++
				}
				cats += "	</div> " + new_line
			}

			index_level_2++
		}

		cats += "	</div> " + new_line
		index_level_1++
	}

	cats += "</div>" + new_line

	return cats
}

// Генерирует список тега select
// ### надо бы добавить опции
func (c *Categories) RenderSelectList() string {

	var html_select string

	html_select = ""
	new_line := "\r\n"

	// Сортировка по "весу"
	sort.Slice(c.List, func(i, j int) bool {
		if c.List[i].Weight < c.List[j].Weight {
			return true
		}
		if c.List[i].Weight > c.List[j].Weight {
			return false
		}
		return true
	})

	for i := 0; i < len(c.List); i++ {
		if c.List[i].Level != 1 {
			continue
		}

		html_select += "	<option value=\"" + strconv.Itoa(c.List[i].Id) + "\" >"
		html_select += c.List[i].Name
		html_select += "	</option>"

		for j := 0; j < len(c.List); j++ {
			if c.List[j].Level != 2 {
				continue
			} else if c.List[j].ParentId != c.List[i].Id {
				continue
			}

			html_select += "	<option value=\"" + strconv.Itoa(c.List[j].Id) + "\" >"
			html_select += "-- " + c.List[j].Name
			html_select += "	</option>"

			for k := 0; k < len(c.List); k++ {
				if c.List[k].Level != 3 {
					continue
				} else if c.List[k].ParentId != c.List[j].Id {
					continue
				}

				html_select += "	<option value=\"" + strconv.Itoa(c.List[k].Id) + "\" >"
				html_select += "-- -- " + c.List[k].Name
				html_select += "	</option>"
			}
		}
	}

	html_select += new_line

	return html_select
}

// Генерирует список тега select для веса
// просто цифры, ничего с категориями не связано
func (c *Categories) RenderWeightSelectList() string {

	var html_weight_select string
	new_line := "\r\n"

	for i := 1; i < 51; i++ {
		html_weight_select += "	<option value=\"" + strconv.Itoa(i) + "\" >"
		html_weight_select += strconv.Itoa(i)
		html_weight_select += "	</option>"
	}

	html_weight_select += new_line

	return html_weight_select
}

// Проверяет корректность заданных уровней
// и исправляет на те, которые соответствуют иерархии
// по полям id и parent_id
// *
// ### Не реализована ###
func (c *Categories) check_levels() bool {
	return true
}

// Возвращает список подкатегорий
// без вложенных, только прямые потомки
// *
// ### Не реализована ###
func (c *Categories) ChildListOf(Id int) []models.Category {

	var Children []models.Category

	return Children
}

// Возвращает кол-во подкатегорий
// без вложенных, только прямые потомки
// *
func (c *Categories) ChildCountOf(Id int) int {

	count := 0

	for i := 0; i < len(c.List); i++ {
		if c.List[i].ParentId == Id {
			count++
		}
	}

	return count
}
