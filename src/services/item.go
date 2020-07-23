package services

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/database/models"
	"github.com/nupamore/Collection-Server/src/services/response"
)

// ItemService : Items services
type ItemService struct{}

// GetUsersItems : GetUsersItems
// @Summary 특정 유저가 가진 모든 아이템들을 가져온다
// @Param memberID path int true "멤버 아이디"
// @Success 0 {object} response.JSONResult{data=[]models.Item}
// @Router /members/{memberID}/items [get]
func (ItemService) GetUsersItems(c echo.Context) error {
	memberID := c.Param("memberID")
	db, _ := c.Get("db").(*gorm.DB)

	var res response.JSONResult
	var items []models.Item
	db.Where("memberID = ?", memberID).Find(&items)

	res.Data = items
	return c.JSON(http.StatusOK, res)
}
