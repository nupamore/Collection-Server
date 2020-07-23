package services

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/database/models"
	"github.com/nupamore/Collection-Server/src/services/response"
)

// MemberService : Members services
type MemberService struct{}

// GetAllMembers : GetAllMembers
// @Summary 모든 멤버 목록을 불러온다
// @Success 0 {object} response.JSONResult{data=[]models.Member}
// @Router /members [get]
func (MemberService) GetAllMembers(c echo.Context) error {
	db, _ := c.Get("db").(*gorm.DB)
	var res response.JSONResult
	var members []models.Member
	db.Find(&members)

	res.Data = members
	return c.JSON(http.StatusOK, res)
}

// GetMember : GetMember
// @Summary 특정 멤버 정보를 불러온다
// @Param memberID path int true "멤버 아이디"
// @Success 0 {object} response.JSONResult{data=models.Member}
// @Failure 9001
// @Router /members/{memberID} [get]
func (MemberService) GetMember(c echo.Context) error {
	id := c.Param("memberID")
	db, _ := c.Get("db").(*gorm.DB)

	var res response.JSONResult
	var member models.Member

	if err := db.Where("id = ?", id).First(&member).Error; gorm.IsRecordNotFoundError(err) {
		res.Code = response.StatusNotExist
		res.Message = response.StatusText(res.Code)
	} else {
		res.Data = member
	}
	return c.JSON(http.StatusOK, res)
}
