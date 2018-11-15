package main

import (
	"cat/conf"
	"cat/controllers"

	"cat/models"
	ent "cat/models/entities"

	_ "cat/routers"
	"database/sql"

	//	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var App models.Application

func init() {

	var err error

	// Инициализация подключения в БД PostgreSQL
	db_info := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		conf.PostgresUser, conf.PostgresPassword, conf.PostgresDB)

	conf.DB_postgres, err = sql.Open("postgres", db_info)
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация подключения ORM
	err = orm.RegisterDriver("postgres", orm.DRPostgres)

	if err != nil {
		fmt.Println("init - orm.RegisterDriver Error: ", err)
		// Не запускать
		return
	}

	maxIdle := 30
	maxConn := 30

	orm_info := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable",
		conf.PostgresUser, conf.PostgresPassword, conf.PostgresDB)

	err = orm.RegisterDataBase("default", "postgres", orm_info, maxIdle, maxConn)

	if err != nil {
		fmt.Println("init - orm.RegisterDataBase Error: ", err)
		// Не запускать
		return
	}

	// Регистрация моделей в ORM
	orm.RegisterModel(new(ent.Category))

	conf.AppOrm = orm.NewOrm()
	conf.AppOrm.Using("default")

	// Подготовка данных приложения
	App.Categories.Load()
	controllers.App = &App

}

func main() {

	cat := ent.Category{Id: 5}

	err := conf.AppOrm.Read(&cat)

	fmt.Println(err)

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(cat)
	}

	fmt.Println("start")
	//return

	for i := 0; i < len(App.Categories.List); i++ {
		fmt.Println(*App.Categories.List[i])
	}

	//App.Categories.Add(5, 3, 12, "Ещё категория 2", false, "/")

	//App.Categories.List[1].Name += "_UPDATED"
	//App.Categories.List[1].Update()
	fmt.Println("----------------")

	for i := 0; i < len(App.Categories.List); i++ {
		fmt.Println(*App.Categories.List[i])
	}

	/*
		fmt.Println("sort")
		fmt.Println(App.Categories.RenderToAdminPanel())

		for i := 0; i < len(App.Categories.List); i++ {
			fmt.Println(App.Categories.List[i])
		}

		fmt.Println("end")

		fmt.Println("start beego!")
	*/
	//beego.Run()

}
