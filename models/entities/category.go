package models

import (
	"cat/conf"
	"strconv"

	//"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Category struct {
	Id int

	ParentId int // Идентификатор родительской категории

	// Уровень вложенности категории
	// вычисляемый параметр на основе поля ParentId
	// введён для удобства (категории меняются редко)
	// иерархия фиксируется вручную, а не вычисляется каждый раз
	Level int

	Weight int    // Вес категории для сортировки
	Name   string // Название категории
	Active bool   // Активна/отображается или нет
	URL    string // Адрес категории

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

	query := "INSERT INTO public.category( "
	query += "  id, parent_id, level, weight, name, active, url)"
	query += "VALUES ( "
	query += string(max_id) + ", "
	query += strconv.Itoa(c.ParentId) + ", "
	query += strconv.Itoa(c.Level) + ", "
	query += strconv.Itoa(c.Weight) + ", "
	query += "'" + c.Name + "', "
	query += " false ,"
	query += "'" + c.URL + "' "
	query += ");"

	_, err = conf.DB_postgres.Exec(query)

	if err != nil {
		//fmt.Println(query)
		log.Fatal(err)
		return false
	}

	return true
}

func (c *Category) Update() bool {

	query := "UPDATE public.category "
	query += "SET "
	query += "  parent_id = " + strconv.Itoa(c.ParentId) + " , "
	query += "  level = " + strconv.Itoa(c.Level) + ", "
	query += "  weight = " + strconv.Itoa(c.Weight) + ", "
	query += "  name = '" + c.Name + "', "
	if c.Active == true {
		query += "  active = true, "
	} else {
		query += "  active = false, "
	}
	query += "  name = '" + c.URL + "' "
	query += "WHERE "
	query += "  id = " + strconv.Itoa(c.Id) + ";"

	_, err := conf.DB_postgres.Exec(query)

	if err != nil {
		//fmt.Println(sql_res)
		log.Fatal(err)
		return false
	}

	return true
}

func (c *Category) Delete() bool {

	query := "DELETE FROM public.category "
	query += "WHERE "
	query += "  id = " + strconv.Itoa(c.Id) + ";"

	_, err := conf.DB_postgres.Exec(query)

	if err != nil {
		//fmt.Println(sql_res)
		log.Fatal(err)
		return false
	}

	return true
}
