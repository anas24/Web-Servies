package main

import (
	"database/sql"
	"mvc-rest/controller"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	initDB()
	controller.HandleRequest()
}

func initDB() {
	db, err = sql.Open("mysql", "root:qwaszx12@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error)
	}
}
