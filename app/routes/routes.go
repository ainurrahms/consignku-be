package routes

import (
	"consignku/controller/discounts"
	"consignku/controller/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteLists struct {
	JWTMiddleware       middleware.JWTConfig
	UserController      users.UserController
	DiscountsController discounts.DiscountsController
}

func (r *RouteLists) RouteRegister(e *echo.Echo) {
	users := e.Group("v1/api/auth")
	users.POST("/register", r.UserController.Register)
	users.POST("/login", r.UserController.Login)

	discounts := e.Group("v1/api/discounts")
	//get all belum :(
	discounts.GET("/all", r.DiscountsController.Fetch, middleware.JWTWithConfig(r.JWTMiddleware))

	discounts.POST("/create", r.DiscountsController.Store, middleware.JWTWithConfig(r.JWTMiddleware))
	discounts.GET("/id/:id", r.DiscountsController.GetByID, middleware.JWTWithConfig(r.JWTMiddleware))
	discounts.PUT("/update/:id", r.DiscountsController.Update, middleware.JWTWithConfig(r.JWTMiddleware))
	discounts.DELETE("/delete/:id", r.DiscountsController.Delete, middleware.JWTWithConfig(r.JWTMiddleware))
}
