package server

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// глобальная переменная с подключением к БД
var Db *sqlx.DB

// функция, инициирующая подключение к БД
func InitDB() (err error) {
	//строка, содержащая данные для подключения к БД в следующем формате:
	//login:password@tcp(host:port)/dbname
	var dataSourceName = "web:6iqvXXX@tcp(remotemysql.com:3306)/0vfkXXX"
	//подключаемся к БД, используя нужный драйвер и данные для подключения
	Db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		return
	}
	err = Db.Ping()
	return
}
