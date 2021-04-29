package routes

import (
	"project/controllers"
	mid "project/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	mid.LogMiddleware(e)
	e.GET("/product", controllers.GetProductController)
	jwt := e.Group("")
	jwt.POST("/product", controllers.CreateProductController)

	return e
}
