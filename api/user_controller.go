package api

import (
	"github.com/labstack/echo/v4"
	"go-campaign-no-02/internal/service"
)

type UserController interface {
	GetUser(c echo.Context) error
	GetUserById(c echo.Context) error
	CreateUser(c echo.Context) error
}
type UserControllerImpl struct {
	userService service.UserService
}

var instanceUserController *UserControllerImpl
