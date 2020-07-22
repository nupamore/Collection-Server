package routes

import "github.com/labstack/echo"

type router struct{}

// Routes : Init Routes
func Routes(g *echo.Group) {
	var r router
	g.GET("/members/:id", r.getMember)
	g.GET("/items/:id", r.getItem)
}
