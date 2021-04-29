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
<<<<<<< HEAD
	e.GET("/product", controllers.GetProductController)
	jwt := e.Group("")
	jwt.POST("/product", controllers.CreateProductController)
	jwt.POST("/cart", controllers.CreateCartController)
	jwt.GET("/cart", controllers.GetCartController)
	jwt.DELETE("/cart/:id", controllers.DeleteCartController)
	jwt.PUT("/cart/:id", controllers.DeleteCartController)

=======
>>>>>>> user
	return e
}
