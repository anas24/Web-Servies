package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

func InitDB() {
	Db, err = sql.Open("mysql", "root:qwaszx12@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error)
	}
	fmt.Println("DB Object created")
	fmt.Println(Db)
}
