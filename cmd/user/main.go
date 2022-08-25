package main

import (
	"github.com/labstack/echo/v4"
	"go-campaign-no-02/api"
	"go-campaign-no-02/db"
)

func main() {
	db.Init()
	e := echo.New()
	api.UserControllerRouter(e.Group(""))
	e.Logger.Fatal(e.Start(":9090"))
}
