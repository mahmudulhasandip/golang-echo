package main

import (
	"echo-auth/pkg/config"
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
	routes.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
