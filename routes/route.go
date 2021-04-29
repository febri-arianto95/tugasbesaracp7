package routes

import (
	"project/controllers"
	mid "project/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	mid.LogMiddleware(e)
	e.POST("/users", controllers.CreateUsersController)
	e.POST("/login", controllers.LoginUsersController)
	return e
}
