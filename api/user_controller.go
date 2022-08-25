package api

import (
	"github.com/labstack/echo/v4"
	"go-campaign-no-02/internal/dto"
	"go-campaign-no-02/internal/service"
	"net/http"
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

func NewUserController() *UserControllerImpl {
	if instanceUserController == nil {
		instanceUserController = &UserControllerImpl{
			userService: service.NewUserService(),
		}
	}
	return instanceUserController
}

func (r UserControllerImpl) GetUser(c echo.Context) error {
	listUser := r.userService.GetUser()
	return c.JSON(http.StatusOK, listUser)
}

func (r UserControllerImpl) GetUserById(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r UserControllerImpl) CreateUser(c echo.Context) error {
	userReq := dto.UserReq{}
	err := c.Bind(&userReq)
	if err != nil {
		return err
	}

	user := r.userService.CreateUser(userReq.Name, userReq.Deps, userReq.Age)

	return c.JSON(http.StatusOK, user)
}

func UserControllerRouter(e *echo.Group) {
	controller := NewUserController()
	e.GET("/all-user", controller.GetUser)
	e.POST("/create", controller.CreateUser)
}
