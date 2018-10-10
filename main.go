package main

import (
	"cat/conf"
	"cat/controllers"
	"cat/models"

	_ "cat/routers"
	"database/sql"

	"github.com/astaxie/beego"

	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var App models.Application

func init() {

	var err error

	// Инициализация подключения в БД PostgreSQL
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		conf.PostgresUser, conf.PostgresPassword, conf.PostgresDB)

	conf.DB_postgres, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database connected")
		App.Categories.Load()
		controllers.App = &App
	}

}

func main() {

	/*
		var app app.Application

		app.Categories.Load()

		app.Categories.Remove(2)
		app.Categories.Remove(3)

		app.Categories.Add(0, 1, 15, "Первая")
		app.Categories.Add(0, 1, 16, "Вторая")
		app.Categories.Add(0, 1, 16, "Третья")

		fmt.Println("start")

		for i := 0; i < len(app.Categories.List); i++ {
			fmt.Println(app.Categories.List[i])
		}

		fmt.Println("end")

		return
	*/

	beego.Run()
}
