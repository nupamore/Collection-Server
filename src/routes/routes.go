package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/controllers"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Routes : Init Routes
func Routes(g *echo.Group) {
	g.GET("/swagger/*", echoSwagger.WrapHandler)

	var m controllers.Member
	g.GET("/members", m.GetAllMembers)
	g.POST("/members", m.CreateMember)
	g.GET("/members/:memberId", m.GetMember)
	g.PUT("/members/:memberId", m.UpdateMember)
	g.DELETE("/members/:memberId", m.DeleteMember)

	var i controllers.Item
	g.GET("/members/:memberId/items", i.GetUsersItems)
	g.POST("/members/:memberId/items", i.AddUsersItems)
	g.DELETE("/members/:memberId/items", i.SubtractUsersItems)

	var c controllers.Char
	g.GET("/members/:memberId/characters", c.GetUsersCharacters)
	g.POST("/members/:memberId/characters", c.CreateUsersCharacters)
	g.PUT("/members/:memberId/characters", c.UpdateUsersCharacters)
	g.DELETE("/members/:memberId/characters", c.DeleteUsersCharacters)

	var e controllers.Equip
	g.GET("/members/:memberId/equipments", e.GetUsersEquipments)
	g.POST("/members/:memberId/equipments", e.CreateUsersEquipments)
	g.PUT("/members/:memberId/equipments", e.UpdateUsersEquipments)
	g.DELETE("/members/:memberId/equipments", e.DeleteUsersEquipments)

	var q controllers.Quest
	g.GET("/members/:memberId/quests", q.GetUsersQuests)
	g.POST("/members/:memberId/quests", q.CreateUsersQuests)
	g.PUT("/members/:memberId/quests", q.UpdateUsersQuests)
	g.DELETE("/members/:memberId/quests", q.DeleteUsersQuests)
}
