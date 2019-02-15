package main

import (
	"mvc-rest/controller"
	"mvc-rest/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.InitDB()
	controller.HandleRequest()
}
