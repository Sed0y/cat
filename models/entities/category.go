package models

import (
	"cat/conf"

	"strconv"

	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func init() {

}

type Category struct {
	Id int `form:"-"`

	// Название категории
	Name string `form:"name,text,name:" valid:"MinSize(5);MaxSize(30)"`

	ParentId int `form:"parent_id"` // Идентификатор родительской категории

	// Уровень вложенности категории
	// вычисляемый параметр на основе поля ParentId
	// введён для удобства (категории меняются редко)
	// иерархия фиксируется вручную, а не вычисляется каждый раз
	Level int `form:"-"`

	Weight int `form:"wheight"` // Вес категории для сортировки

	Active bool   `form:"active"` // Активна/отображается или нет
	Url    string `form:"url"`    // Адрес категории

}

// Название таблицы в БД
func (c *Category) TableName() string {
	return "category"
}

func (c *Category) Create() bool {

	var max_id []byte
	var level []byte

	rows, err := conf.DB_postgres.Query("SELECT CASE WHEN max(id) is NULL THEN 1 ELSE max(id) + 1  END FROM public.category;")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
		return false
	}

	for rows.Next() {
		err = rows.Scan(&max_id)

		if err != nil {
			log.Fatal(err)
			return false
		}
	}

	c.Id, _ = strconv.Atoi(string(max_id))

	fmt.Println(c.ParentId)

	if c.ParentId == 0 {
		c.Level = 1
	} else {
		rows, err = conf.DB_postgres.Query("SELECT level + 1 FROM public.category WHERE id = " + strconv.Itoa(c.ParentId) + ";")
		if err != nil {
			log.Fatal(err)
			return false
		}

		for rows.Next() {
			err = rows.Scan(&level)

			if err != nil {
				log.Fatal(err)
				return false
			}
		}

		c.Level, _ = strconv.Atoi(string(level))
	}

	_, err = conf.AppOrm.Insert(c)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true

}

func (c *Category) Update() bool {

	conf.AppOrm.Update(c)
	return true
}

func (c *Category) Delete() bool {

	conf.AppOrm.Delete(c)
	return true
}
