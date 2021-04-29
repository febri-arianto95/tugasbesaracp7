package routes

import (
	"project/controllers"
	mid "project/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	mid.LogMiddleware(e)
	e.POST("/product", controllers.CreateProductController)
	e.GET("/product", controllers.GetProductController)

	return e
}
