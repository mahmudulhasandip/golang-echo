package routes

import (
	"echo-auth/pkg/auth"
	"echo-auth/pkg/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var RegisterUserRoutes = func(router *echo.Group) {
	router.POST("/signup", controllers.SignupController).Name = "signup"
	router.POST("/login", controllers.LoginController).Name = "login"

	// jwt authentication middleware
	router.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &auth.Claims{},
		SigningKey:              []byte(auth.GetJWTSecret()),
		TokenLookup:             "cookie:access-token", // "<source>:<name>"
		ErrorHandlerWithContext: auth.JWTErrorChecker,
	}))
	// refresh token middleware
	router.Use(auth.TokenRefresherMiddleware)

	router.GET("/user", controllers.UserController).Name = "user"
}
