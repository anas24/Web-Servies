package model

import (
	"fmt"
	"mvc-rest/database"
	"net/http"
	"strconv"
)

type User struct {
	UserId int    `json: "id"`
	Name   string `json:"name"`
}

const ErrorMessage string = "Something went wrong"
const InsertionSuccessMessage string = "Successfully Inserted"
const UpdationSuccessMessage string = "Successfully Updated"
const DeletionSuccessMessage string = "Successfully Deleted"
const InvalidParamsMessage string = "It's not funny to give invalid params"

func AllUsers(w http.ResponseWriter, r *http.Request) interface{} {
	var users []User
	pagNo := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("size")
	pagNoInt, er1 := strconv.Atoi(pagNo)
	pageSizeInt, er2 := strconv.Atoi(pageSize)
	if er1 == nil && er2 == nil {

		allUsers, err := database.Db.Prepare("Select id, name From users limit ?, ?")
		defer allUsers.Close()
		left := strconv.Itoa((pagNoInt - 1) * pageSizeInt)
		right := strconv.Itoa(pageSizeInt)
		if err != nil {
			fmt.Println(err)
			return ErrorMessage
		}
		fmt.Println(left, right)
		allUsers1, err := allUsers.Query(left, right)
		if err != nil {
			fmt.Println(err)
			return ErrorMessage
		}
		for allUsers1.Next() {
			var id int
			var name string
			err := allUsers1.Scan(&id, &name)
			if err != nil {
				fmt.Println(err)
			} else {
				userData := User{id, name}
				users = append(users, userData)
				fmt.Println(id, name)
			}
		}
		return users
	} else {
		return InvalidParamsMessage
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) interface{} {
	nId := r.URL.Query().Get("id")
	user, err := database.Db.Query("SELECT id, name FROM users where id = ?", nId)
	defer user.Close()
	if err != nil {
		fmt.Println(err)
		return ErrorMessage
	}
	if user != nil {
		for user.Next() {
			var id int
			var name string
			err := user.Scan(&id, &name)
			if err != nil {
				panic(err.Error())
			} else {
				userData := User{id, name}
				fmt.Println(id, name)
				return userData
			}
		}
	}
	return ErrorMessage
}

func Insert(w http.ResponseWriter, r *http.Request) bool {
	nId := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	newUser, err := database.Db.Prepare("Insert into users(id, name) values(?, ?)")
	defer newUser.Close()
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "Params Invalid")
		return false
	}
	if newUser != nil {
		newUser.Exec(nId, name)
		fmt.Println("Record inserted")
		return true
	}
	return false
}
func Update(w http.ResponseWriter, r *http.Request) bool {
	nId := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	updated, err := database.Db.Prepare("Update users set name=? where id = ?")
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "Params Invalid")
		return false
	}

	if updated != nil {
		updated.Exec(name, nId)
		fmt.Println("Record updated")
		return true
	}
	return false
}
func Delete(w http.ResponseWriter, r *http.Request) bool {
	nId := r.URL.Query().Get("id")
	deleted, err := database.Db.Prepare("Delete from users where id=?")
	if err != nil {
		fmt.Println(err)
		return false
	}

	if deleted != nil {
		deleted.Exec(nId)
		fmt.Println("Record deleted")
		return true
	}
	return false
}
