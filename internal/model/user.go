package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   int `gorm:"primaryKey"`
	Name string
	Deps string
	Age  int
}

func NewUser(name string, deps string, age int) *User {
	return &User{
		Name: name,
		Deps: deps,
		Age:  age,
	}
}
