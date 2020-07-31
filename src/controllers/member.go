package controllers

import (
	"net/http"

	"github.com/fatih/structs"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/controllers/response"
	"github.com/nupamore/Collection-Server/src/database"
	"github.com/nupamore/Collection-Server/src/database/models"
	"github.com/nupamore/Collection-Server/src/services"
)

// Member : Members controller
type Member struct{}

// GetAllMembers : GetAllMembers
// @Summary 모든 멤버 목록을 불러온다
// @Success 0 {object} response.JSONResult{data=[]models.Member}
// @Router /members [get]
func (Member) GetAllMembers(c echo.Context) error {
	var res response.JSONResult
	var members []models.Member
	db, _ := c.Get("db").(*gorm.DB)

	db.Find(&members)
	res.Data = members

	return c.JSON(http.StatusOK, res)
}

// CreateMember : CreateMember
// @Summary 멤버를 생성한다
// @Param body body models.Member true "멤버 모델"
// @Success 0
// @Router /members [post]
func (Member) CreateMember(c echo.Context) error {
	var res response.JSONResult
	var member models.Member
	db, _ := c.Get("db").(*gorm.DB)

	c.Bind(&member)
	if err := db.Create(&member).Error; err != nil {
		if notExist := db.NewRecord(member); !notExist {
			res.Code = response.StatusExist
			res.Message = response.StatusText(res.Code)
		} else {
			code, message := database.ParseError(err)
			res.Code = code
			res.Message = message
		}
	} else {
		res.Message = response.StatusText(response.StatusCreated)
	}

	return c.JSON(http.StatusOK, res)
}

// GetMember : GetMember
// @Summary 특정 멤버 정보를 불러온다
// @Param memberId path string true "멤버 아이디"
// @Success 0 {object} response.JSONResult{data=models.Member}
// @Failure 9011
// @Router /members/{memberId} [get]
func (Member) GetMember(c echo.Context) error {
	var res response.JSONResult
	id := c.Param("memberId")

	member, err := services.ValidMember(id)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	res.Data = member

	return c.JSON(http.StatusOK, res)

}

// UpdateMember : UpdateMember
// @Summary 특정 멤버를 수정한다
// @Param memberId path string true "멤버 아이디"
// @Param body body models.Member true "멤버 모델"
// @Success 0
// @Failure 9011
// @Router /members/{memberId} [put]
func (Member) UpdateMember(c echo.Context) error {
	var res response.JSONResult
	db, _ := c.Get("db").(*gorm.DB)
	id := c.Param("memberId")

	before, err := services.ValidMember(id)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}

	var after models.Member
	c.Bind(&after)
	db.Model(&before).Update(structs.Map(after))
	res.Message = response.StatusText(response.StatusUpdated)

	return c.JSON(http.StatusOK, res)
}

// DeleteMember : DeleteMember
// @Summary 특정 멤버를 삭제한다
// @Param memberId path string true "멤버 아이디"
// @Success 0
// @Failure 9011
// @Router /members/{memberId} [delete]
func (Member) DeleteMember(c echo.Context) error {
	var res response.JSONResult
	db, _ := c.Get("db").(*gorm.DB)
	id := c.Param("memberId")

	member, err := services.ValidMember(id)
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	db.Delete(member)
	res.Message = response.StatusText(response.StatusDeleted)

	return c.JSON(http.StatusOK, res)
}
