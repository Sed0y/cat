package conf

import (
	"database/sql"
)

var DB_postgres *sql.DB

const (

	// параметры подключения к БД PostgreSQL
	PostgresUser     = "postgres"
	PostgresPassword = "vmmjn28z"
	PostgresDB       = "ctlg"

	// *************  НАСТРОЙКИ КЛИЕНТА   ********************************
	// хранение куки
	CookieExpiration = 180 // дней

)
