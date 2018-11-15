package conf

import (
	"database/sql"

	"github.com/astaxie/beego/orm"
)

var DB_postgres *sql.DB
var AppOrm orm.Ormer

const (

	// параметры подключения к БД PostgreSQL
	PostgresUser     = "postgres"
	PostgresPassword = "vmmjn28z"
	PostgresDB       = "ctlg"

	// *************  НАСТРОЙКИ КЛИЕНТА   ********************************
	// хранение куки
	CookieExpiration = 180 // дней

)
