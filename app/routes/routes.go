package routes

import (
	"consignku/controller/discounts"
	"consignku/controller/product_types"
	"consignku/controller/product_used_times"
	"consignku/controller/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteLists struct {
	JWTMiddleware              middleware.JWTConfig
	UserController             users.UserController
	DiscountsController        discounts.DiscountsController
	ProductTypesController     product_types.ProductTypesController
	ProductUsedTimesController product_used_times.ProductUsedTimesController
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

	productType := e.Group("v1/api/product/type")
	//get all belum :(
	// discounts.GET("/all", r.DiscountsController.Fetch, middleware.JWTWithConfig(r.JWTMiddleware))

	productType.POST("/create", r.ProductTypesController.Store, middleware.JWTWithConfig(r.JWTMiddleware))
	productType.GET("/id/:id", r.ProductTypesController.GetByID, middleware.JWTWithConfig(r.JWTMiddleware))
	productType.PUT("/update/:id", r.ProductTypesController.Update, middleware.JWTWithConfig(r.JWTMiddleware))
	productType.DELETE("/delete/:id", r.ProductTypesController.Delete, middleware.JWTWithConfig(r.JWTMiddleware))

	productUsedTimes := e.Group("v1/api/product/used")
	//get all belum :(
	// discounts.GET("/all", r.DiscountsController.Fetch, middleware.JWTWithConfig(r.JWTMiddleware))

	productUsedTimes.POST("/create", r.ProductUsedTimesController.Store, middleware.JWTWithConfig(r.JWTMiddleware))
	productUsedTimes.GET("/id/:id", r.ProductUsedTimesController.GetByID, middleware.JWTWithConfig(r.JWTMiddleware))
	productUsedTimes.PUT("/update/:id", r.ProductUsedTimesController.Update, middleware.JWTWithConfig(r.JWTMiddleware))
	productUsedTimes.DELETE("/delete/:id", r.ProductUsedTimesController.Delete, middleware.JWTWithConfig(r.JWTMiddleware))

}
