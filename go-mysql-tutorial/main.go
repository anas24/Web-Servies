package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:qwaszx12@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	fmt.Println("Successfully Connected to Mysql database")

	insert, err := db.Query("INSERT INTO users VALUES(6, 'Qazi2');")

	if err != nil {
		panic(err.Error)
	}
	defer insert.Close()

	fmt.Println("Successfully Inserted")

	allRecords, err := db.Query("SELECT * FROM users;")

	if err != nil {
		panic(err.Error)
	} else {
		for allRecords.Next() {
			var id int
			var name string
			err := allRecords.Scan(&id, &name)
			if err != nil {
				panic(err)
			} else {
				fmt.Println(id, "  ", name)
			}
		}
	}
	defer allRecords.Close()
}
