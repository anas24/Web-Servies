package controller

import (
	"encoding/json"
	"fmt"
	"mvc-rest/database"
	_ "mvc-rest/database"
	"mvc-rest/model"
	"net/http"
)

var users []model.User

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(database.Db)

	allUsers, err := database.Db.Query("SELECT id, name FROM users")
	defer allUsers.Close()
	defer database.Db.Close()
	if err != nil {
		fmt.Println(err)
	}

	for allUsers.Next() {
		var id int
		var name string
		err := allUsers.Scan(&id, &name)
		if err != nil {
			panic(err)
		} else {
			fmt.Println(id, "  ", name)
		}
		user := model.User{id, name}
		users = append(users, user)
	}
	fmt.Println("EndPoint hit: All articles endpoint", allUsers, err)
	json.NewEncoder(w).Encode(users)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page endPoinHit")
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	nId := r.URL.Query().Get("id")
	user, err := database.Db.Query("SELECT id, name FROM users where id = ?", nId)
	defer user.Close()
	if err != nil {
		fmt.Println(err)
	}
	if user != nil {
		for user.Next() {
			var id int
			var name string
			err := user.Scan(&id, &name)
			if err != nil {
				panic(err.Error())
			} else {
				userData := model.User{id, name}
				fmt.Println(id, name)
				json.NewEncoder(w).Encode(userData)
			}
		}
	}
}
