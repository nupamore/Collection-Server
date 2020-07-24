package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/controllers/response"
	"github.com/nupamore/Collection-Server/src/database"
	"github.com/nupamore/Collection-Server/src/database/models"
)

// MemberCtrl : Members controller
type MemberCtrl struct{}

// GetAllMembers : GetAllMembers
// @Summary 모든 멤버 목록을 불러온다
// @Success 0 {object} response.JSONResult{data=[]models.Member}
// @Router /members [get]
func (MemberCtrl) GetAllMembers(c echo.Context) error {
	var res response.JSONResult
	var members []models.Member
	db, _ := c.Get("db").(*gorm.DB)

	db.Find(&members)
	res.Data = members

	return c.JSON(http.StatusOK, res)
}

// CreateMember : CreateMember
// @Summary 멤버를 생성한다
// @Param json body models.Member true "멤버 모델"
// @Success 0
// @Router /members [post]
func (MemberCtrl) CreateMember(c echo.Context) error {
	var res response.JSONResult
	var member models.Member
	db, _ := c.Get("db").(*gorm.DB)

	c.Bind(&member)
	if err := db.Create(member).Error; err != nil {
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
// @Param memberID path string true "멤버 아이디"
// @Success 0 {object} response.JSONResult{data=models.Member}
// @Failure 9001
// @Router /members/{memberID} [get]
func (MemberCtrl) GetMember(c echo.Context) error {
	var res response.JSONResult
	var member models.Member
	id := c.Param("memberID")
	db, _ := c.Get("db").(*gorm.DB)

	if err := db.Where("id = ?", id).First(&member).Error; gorm.IsRecordNotFoundError(err) {
		res.Code = response.StatusNotExist
		res.Message = response.StatusText(res.Code)
	} else {
		res.Data = member
	}

	return c.JSON(http.StatusOK, res)
}

// UpdateMember : UpdateMember
// @Summary 특정 멤버를 수정한다
// @Param memberID path string true "멤버 아이디"
// @Param json body models.Member true "멤버 모델"
// @Success 0
// @Failure 9001
// @Router /members/{memberID} [put]
func (MemberCtrl) UpdateMember(c echo.Context) error {
	var res response.JSONResult
	var before, after models.Member
	db, _ := c.Get("db").(*gorm.DB)
	id := c.Param("memberID")

	if err := db.Where("id = ?", id).First(&before).Error; gorm.IsRecordNotFoundError(err) {
		res.Code = response.StatusNotExist
		res.Message = response.StatusText(res.Code)
	} else {
		c.Bind(&after)
		db.Model(&before).Update(after)
		res.Message = response.StatusText(response.StatusUpdated)
	}

	return c.JSON(http.StatusOK, res)
}

// DeleteMember : DeleteMember
// @Summary 특정 멤버를 삭제한다
// @Param memberID path string true "멤버 아이디"
// @Success 0
// @Failure 9001
// @Router /members/{memberID} [delete]
func (MemberCtrl) DeleteMember(c echo.Context) error {
	var res response.JSONResult
	var member models.Member
	db, _ := c.Get("db").(*gorm.DB)
	id := c.Param("memberID")

	if err := db.Where("id = ?", id).First(&member).Error; gorm.IsRecordNotFoundError(err) {
		res.Code = response.StatusNotExist
		res.Message = response.StatusText(res.Code)
	} else {
		db.Delete(member)
		res.Message = response.StatusText(response.StatusDeleted)
	}

	return c.JSON(http.StatusOK, res)
}
