package routes

import (
	"github.com/labstack/echo"
	"github.com/nupamore/Collection-Server/src/services"
)

// Routes : Init Routes
func Routes(g *echo.Group) {
	var m services.MemberService
	g.GET("/members", m.GetAllMembers)
	g.GET("/members/:id", m.GetMember)

	var i services.ItemService
	g.GET("/items/:id", i.GetItem)
}
