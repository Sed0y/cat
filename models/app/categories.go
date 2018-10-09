package app

import (
	"cat/conf"
	"cat/models/entities"
	//"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type Categories struct {
	List []models.Category
}

func (c *Categories) Load() bool {

	var id []byte
	var parent_id []byte
	var level []byte
	var weight []byte
	var name []byte

	var current models.Category

	rows, err := conf.DB_postgres.Query("SELECT id, parent_id, level, weight, name FROM public.category;")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
		return false
	}

	c.List = c.List[:0]

	for rows.Next() {
		err = rows.Scan(&id, &parent_id, &level, &weight, &name)

		if err != nil {
			log.Fatal(err)
			return false
		}

		current.Id, _ = strconv.Atoi(string(id))
		current.ParentId, _ = strconv.Atoi(string(parent_id))
		current.Level, _ = strconv.Atoi(string(level))
		current.Weight, _ = strconv.Atoi(string(weight))
		current.Name = string(name)

		c.List = append(c.List, current)
	}

	return true
}

func (c *Categories) Add(ParentId int, Level int, Weight int, Name string) bool {

	var NewOne models.Category

	NewOne.ParentId = ParentId
	NewOne.Level = Level
	NewOne.Weight = Weight
	NewOne.Name = Name

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

func (c *Categories) Remove(CategoryID int) bool {

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
