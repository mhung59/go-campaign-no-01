package service

import (
	"go-campaign-no-02/db"
	"go-campaign-no-02/internal/model"
	"gorm.io/gorm"
)

type UserService interface {
	UpdateInfoUser(id int, values string)
	CreateUser(id int, name string, deps string, age int) model.User
	GetUser() []model.User
	GetUserById(id int) model.User
}

type UserServiceImpl struct {
	userDb *gorm.DB
}

func UpdateInfoUser(userModify model.User) {
	//
}

func CreateUser(name string, deps string, age int) model.User {
	u := model.NewUser(name, deps, age)
	db.Manager().Create(u)
	return *u
}

func GetUser() []model.User {
	var users []model.User
	db.Manager().Find(&users)

	return users
}
