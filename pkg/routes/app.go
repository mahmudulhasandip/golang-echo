package routes

import "github.com/labstack/echo/v4"

var RegisterRoutes = func(route *echo.Echo) {
	//route Group
	apiV1 := route.Group("/api/v1")
	// user route Group register
	userGroup := apiV1.Group("/user")
	RegisterUserRoutes(userGroup)

}
