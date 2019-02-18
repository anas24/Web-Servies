package controller

import (
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users", AllUsers)
	http.HandleFunc("/users/{id}", GetUser)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
