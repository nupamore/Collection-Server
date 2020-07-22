package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/nupamore/Collection-Server/src/database"
	"github.com/nupamore/Collection-Server/src/routes"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	db := database.Connect()
	defer db.Close()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(database.ContextDB(db))

	// Routes
	routes.Routes(e.Group(""))

	// Start server
	e.Logger.Fatal(e.Start(":9001"))
}
