package controller

import (
	"encoding/json"
	"fmt"
	_ "mvc-rest/database"
	"mvc-rest/model"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.AllUsers(w, r))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page endPoinHit")
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.GetUser(w, r))
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if model.Insert(w, r) {
		json.NewEncoder(w).Encode(model.InsertionSuccessMessage)
	} else {
		json.NewEncoder(w).Encode(model.ErrorMessage)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if model.Update(w, r) {
		json.NewEncoder(w).Encode(model.UpdationSuccessMessage)
	} else {
		json.NewEncoder(w).Encode(model.ErrorMessage)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if model.Delete(w, r) {
		json.NewEncoder(w).Encode(model.DeletionSuccessMessage)
	} else {
		json.NewEncoder(w).Encode(model.ErrorMessage)
	}
}
