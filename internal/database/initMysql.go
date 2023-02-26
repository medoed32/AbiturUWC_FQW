package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DBmoderator *sql.DB
var DBwebInfo *sql.DB

func InitDBmoderator(login string) (err error) {
	DBmoderator, err = initDB(login)
	return
}
func InitDBWebInfo(login string) (err error) {
	DBwebInfo, err = initDB(login)
	return
}

func initDB(login string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", login)
	if err != nil {
		log.Println(err)
	}
	return
}
