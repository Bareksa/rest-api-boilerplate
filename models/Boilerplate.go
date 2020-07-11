package models

type User struct {
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	Age int `json:"age"`
}
