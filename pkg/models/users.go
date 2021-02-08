package models

import (
	"github.com/gin-gonic/gin"
	user "problem/1/build/gen/api"
)

type User struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
	City string `json:"city"`
}

func (u *User) ToProtoBuf() *user.User {
	return &user.User{
		Id:   u.ID,
		Name: u.Name,
		Age:  u.Age,
		City: u.City,
	}
}

func (u *User) ToGinH() gin.H {
	return gin.H{
		"id":   u.ID,
		"name": u.Name,
		"city": u.City,
		"age":  u.Age,
	}
}

// get all
func UserList() ([]User, error) {
	var users []User
	result := Db.Model(&User{}).Find(&users)
	return users, result.Error
}

// get user information by user id
func FilterByID(ID int32) (*User, error) {
	var user User
	qs := Db.Model(&User{}).Where("id = ?", ID)
	result := qs.First(&user)
	return &user, result.Error
}
