package models

//User data user
type User struct {
	Id       string `json:"uid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Lastname string `json:"lastname"`
	Admin    bool   `json:"admin"`
}

type Users []User
