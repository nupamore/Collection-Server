package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/nupamore/Collection-Server/docs"
	"github.com/nupamore/Collection-Server/src/database"
	"github.com/nupamore/Collection-Server/src/routes"
)

func init() {
	godotenv.Load(".env")
}

// @title Collection API
// @version 1.0
// @host collection.nupa.moe
// @BasePath /api
func main() {
	db := database.Connect()
	defer db.Close()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(database.ContextDB(db))

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	routes.Routes(e.Group(""))

	// Start server
	e.Logger.Fatal(e.Start(":9001"))
}
