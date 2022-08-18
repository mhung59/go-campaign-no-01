package service

import (
	"go-campaign-no-02/db"
	"go-campaign-no-02/internal/model"
)

type UserService interface {
	UpdateInfoUser(id int, values string)
	CreateUser() model.User
	GetUser() []model.User
	GetUserById(id int) model.User
}

func UpdateInfoUser(values string) {
	var user model.User
	db.Manager().Model(&user).Update("Name", values)
}

func CreateUser() model.User {
	u := model.NewUser(1, "User", "He is a developer", 23)
	db.Manager().Create(u)
	return *u
}

func GetUser() []model.User {
	var users []model.User
	db.Manager().Find(&users)

	return users
}
