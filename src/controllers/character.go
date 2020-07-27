package controllers

import (
	"net/http"

	"github.com/nupamore/Collection-Server/src/services"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/controllers/response"
	"github.com/nupamore/Collection-Server/src/database/models"
)

// CharCtrl : Characters controller
type CharCtrl struct{}

// GetUsersCharacters : GetUsersCharacters
// @Summary 특정 유저가 가진 모든 캐릭터들을 가져온다
// @Param memberId path string true "멤버 아이디"
// @Success 0 {object} response.JSONResult{data=[]models.Character}
// @Failure 9001
// @Router /members/{memberId}/characters [get]
func (CharCtrl) GetUsersCharacters(c echo.Context) error {
	var res response.JSONResult
	var chars []models.Character
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.Init(c).ValidMember(memberID); err != nil {
		return err
	}

	db.Where("memberId = ?", memberID).Find(&chars)
	res.Data = chars

	return c.JSON(http.StatusOK, res)
}

// CreateUsersCharacters : CreateUsersCharacters
// @Summary 특정 유저에게 캐릭터들을 할당해준다
// @Param memberId path string true "멤버 아이디"
// @Success 0
// @Failure 9001
// @Router /members/{memberId}/characters [post]
func (CharCtrl) CreateUsersCharacters(c echo.Context) error {
	var res response.JSONResult
	return c.JSON(http.StatusOK, res)
}

// UpdateUsersCharacters : UpdateUsersCharacters
// @Summary 특정 유저가 가진 캐릭터들의 정보를 수정한다
// @Param memberId path string true "멤버 아이디"
// @Success 0
// @Failure 9001
// @Router /members/{memberId}/characters [put]
func (CharCtrl) UpdateUsersCharacters(c echo.Context) error {
	var res response.JSONResult
	return c.JSON(http.StatusOK, res)
}

// DeleteUsersCharacters : DeleteUsersCharacters
// @Summary 특정 유저가 가진 캐릭터들을 삭제한다
// @Param memberId path string true "멤버 아이디"
// @Success 0
// @Failure 9001
// @Router /members/{memberId}/characters [delete]
func (CharCtrl) DeleteUsersCharacters(c echo.Context) error {
	var res response.JSONResult
	return c.JSON(http.StatusOK, res)
}
