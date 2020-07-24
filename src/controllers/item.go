package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/controllers/response"
	"github.com/nupamore/Collection-Server/src/database/models"
)

// ItemCtrl : Items controller
type ItemCtrl struct{}

// GetUsersItems : GetUsersItems
// @Summary 특정 유저가 가진 모든 아이템들을 가져온다
// @Param memberID path int true "멤버 아이디"
// @Success 0 {object} response.JSONResult{data=[]models.Item}
// @Router /members/{memberID}/items [get]
func (ItemCtrl) GetUsersItems(c echo.Context) error {
	var res response.JSONResult
	var items []models.Item
	memberID := c.Param("memberID")
	db, _ := c.Get("db").(*gorm.DB)

	db.Where("memberID = ?", memberID).Find(&items)
	res.Data = items

	return c.JSON(http.StatusOK, res)
}
