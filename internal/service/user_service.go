package service

import (
	"go-campaign-no-02/internal/model"
)

type UserService interface {
	UpdateInfoUser(u *model.User, values string)
	CreateUser() model.User
}

func UpdateInfoUser(u *model.User, values string) {
	u.Name = values
}

func CreateUser() model.User {
	u := new(model.User)
	u.Name = "User"
	u.Deps = "He is a developer"
	u.Age = 23

	return *u
}
