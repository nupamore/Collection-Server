package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/nupamore/Collection-Server/routes"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	routes.Routes(e.Group(""))

	// Start server
	e.Logger.Fatal(e.Start(":9001"))
}
