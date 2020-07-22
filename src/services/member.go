package services

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nupamore/Collection-Server/src/database/models"
)

// MemberService : Members services
type MemberService struct{}

// GetAllMembers : get all member list
func (MemberService) GetAllMembers(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)

	var members []models.Member
	db.Find(&members)

	return c.JSON(http.StatusOK, members)
}

// GetMember : get a member
func (MemberService) GetMember(c echo.Context) error {
	id := c.Param("id")
	db, _ := c.Get("db").(*gorm.DB)

	var member models.Member
	db.Where("id = ?", id).First(&member)

	return c.JSON(http.StatusOK, member)
}
