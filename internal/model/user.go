package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   int
	Name string
	Deps string
	Age  int
}

func NewUser(id int, name string, deps string, age int) *User {
	return &User{
		Id:   id,
		Name: name,
		Deps: deps,
		Age:  age,
	}
}
