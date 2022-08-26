package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   int    `gorm:"primaryKey", json:"id"`
	Name string `json:"name"`
	Deps string `json:"deps"`
	Age  int    `json:"age"`
}

func NewUser(name string, deps string, age int) *User {
	return &User{
		Name: name,
		Deps: deps,
		Age:  age,
	}
}
