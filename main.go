package main

import (
	"echo-auth/pkg/config"
	"echo-auth/pkg/controllers"
	"echo-auth/pkg/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	config.LoadEnv()
	config.Connect()
	config.SyncDatabase()
}

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}, latency=${latency_human}\n",
	}))
	// testing route
	e.GET("/test", controllers.TestController).Name = "test_controller"
	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
