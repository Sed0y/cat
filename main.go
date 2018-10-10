package main

import (
	"cat/conf"
	"cat/controllers"
	"cat/models"

	//_ "cat/routers"
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

	fmt.Println("start")

	for i := 0; i < len(App.Categories.List); i++ {
		fmt.Println(App.Categories.List[i])
	}

	fmt.Println("end")

	return

	fmt.Println("start!")
	beego.Run()
}
