package services

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nupamore/Collection-Server/src/database/models"
	"github.com/nupamore/Collection-Server/src/services/response"
)

// MemberService : Members services
type MemberService struct{}

type memberResponse struct {
	response.Response
	Data []models.Member
}

// GetAllMembers : get all member list
func (MemberService) GetAllMembers(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)
	var res memberResponse
	var members []models.Member
	db.Find(&members)

	res.Data = members
	return c.JSON(http.StatusOK, res)
}

// GetMember : get a member
func (MemberService) GetMember(c echo.Context) error {
	id := c.Param("id")
	db, _ := c.Get("db").(*gorm.DB)

	var res memberResponse
	var member models.Member

	if err := db.Where("id = ?", id).First(&member).Error; gorm.IsRecordNotFoundError(err) {
		res.Code = response.StatusNotExist
		res.Message = response.StatusText(res.Code)
	} else {
		res.Data = []models.Member{member}
	}
	return c.JSON(http.StatusOK, res)
}
