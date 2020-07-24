package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/controllers"
)

// Routes : Init Routes
func Routes(g *echo.Group) {
	var m controllers.MemberCtrl
	g.GET("/members", m.GetAllMembers)
	g.POST("/members", m.CreateMember)
	g.GET("/members/:memberID", m.GetMember)
	g.PUT("/members/:memberID", m.UpdateMember)
	g.DELETE("/members/:memberID", m.DeleteMember)

	var i controllers.ItemCtrl
	g.GET("/members/:memberID/items", i.GetUsersItems)
	g.POST("/members/:memberID/items", i.AddUsersItems)
}
