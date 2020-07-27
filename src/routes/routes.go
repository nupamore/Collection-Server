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
	g.GET("/members/:memberId", m.GetMember)
	g.PUT("/members/:memberId", m.UpdateMember)
	g.DELETE("/members/:memberId", m.DeleteMember)

	var i controllers.ItemCtrl
	g.GET("/members/:memberId/items", i.GetUsersItems)
	g.POST("/members/:memberId/items", i.AddUsersItems)
	g.DELETE("/members/:memberId/items", i.SubtractUsersItems)

	var c controllers.CharCtrl
	g.GET("/members/:memberId/characters", c.GetUsersCharacters)
	g.POST("/members/:memberId/characters", c.CreateUsersCharacters)
	g.PUT("/members/:memberId/characters", c.UpdateUsersCharacters)
	g.DELETE("/members/:memberId/characters", c.DeleteUsersCharacters)
}
