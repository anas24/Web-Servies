package model

type User struct {
	UserId int    `json: "id"`
	Name   string `json:"name"`
}
