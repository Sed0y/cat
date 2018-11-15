package models

import (
	"cat/conf"

	"strconv"

	//	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func init() {

}

type Category struct {
	Id int `form:"-"`

	ParentId int // Идентификатор родительской категории

	// Уровень вложенности категории
	// вычисляемый параметр на основе поля ParentId
	// введён для удобства (категории меняются редко)
	// иерархия фиксируется вручную, а не вычисляется каждый раз
	Level int

	Weight int // Вес категории для сортировки

	// Название категории
	Name string `form:"name,text,name:" valid:"MinSize(5);MaxSize(20)"`

	Active bool   // Активна/отображается или нет
	Url    string // Адрес категории

}

// Название таблицы в БД
func (c *Category) TableName() string {
	return "category"
}

func (c *Category) Create() bool {

	var max_id []byte

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
	_, err = conf.AppOrm.Insert(c)
	if err != nil {
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
