package models

type User struct {
	Id int
	Phone string `form:"phone"'`
	Password string `form:"password"`
}
