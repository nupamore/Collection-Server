package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/services"
)

// Routes : Init Routes
func Routes(g *echo.Group) {
	var m services.MemberService
	g.GET("/members", m.GetAllMembers)
	g.GET("/members/:memberID", m.GetMember)

	var i services.ItemService
	g.GET("/members/:memberID/items", i.GetUsersItems)
}
