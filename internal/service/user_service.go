package service

import (
	"go-campaign-no-02/db"
	"go-campaign-no-02/internal/model"
	"gorm.io/gorm"
)

type UserService interface {
	UpdateInfoUser(id int, values string)
	CreateUser(name string, deps string, age int) model.User
	DeleteUser(name string) string
	GetUser() []model.User
	GetUserById(id int) model.User
}

type UserServiceImpl struct {
	dbc *gorm.DB
}

var instanceUserService *UserServiceImpl

func NewUserService() *UserServiceImpl {
	if instanceUserService == nil {
		instanceUserService = &UserServiceImpl{
			dbc: db.Manager(),
		}
	}
	return instanceUserService
}

func (UserServiceImpl) UpdateInfoUser(id int, values string) {
	//TODO implement me
	panic("implement me")
}

func (r UserServiceImpl) CreateUser(name string, deps string, age int) model.User {
	u := model.NewUser(name, deps, age)
	r.dbc.Create(u)
	return *u
}

func (r UserServiceImpl) GetUser() []model.User {
	var users []model.User
	r.dbc.Find(&users)

	return users
}

func (r UserServiceImpl) GetUserById(id int) model.User {
	//TODO implement me
	panic("implement me")
}

func (r UserServiceImpl) DeleteUser(name string) string {
	user := r.dbc.Find("name = ?", name)

	if user == nil {
		return "user not found"
	}

	r.dbc.Delete("name = ?", name)

	return "deleted user " + name
}
