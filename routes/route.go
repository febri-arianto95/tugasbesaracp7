package routes

import (
	"project/constants"
	"project/controllers"
	mid "project/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	mid.LogMiddleware(e)
	e.POST("/users", controllers.CreateUsersController)
	e.POST("/login", controllers.LoginUsersController)
	e.GET("/product", controllers.GetProductController)
	jwt := e.Group("")
	jwt.Use(middleware.JWT([]byte(constants.KEY_JWT)))
	jwt.POST("/product", controllers.CreateProductController)
	jwt.POST("/cart", controllers.CreateCartController)
	jwt.GET("/cart", controllers.GetCartController)
	jwt.DELETE("/cart/:id", controllers.DeleteCartController)
	jwt.PUT("/cart/:id", controllers.UpdateCartController)
	jwt.POST("/checkout", controllers.CreateCheckoutController)
	jwt.GET("/checkout", controllers.GetCheckoutController)
	jwt.POST("/payment", controllers.PaymentController)
	jwt.GET("/payment", controllers.GetPaymentController)

	return e
}
