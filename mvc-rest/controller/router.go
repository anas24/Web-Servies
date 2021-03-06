package controller

import (
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users", AllUsers)
	http.HandleFunc("/get", GetUser)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
