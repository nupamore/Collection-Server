package controllers

import (
	"net/http"

	"github.com/nupamore/Collection-Server/src/database"
	"github.com/nupamore/Collection-Server/src/services"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/controllers/response"
	"github.com/nupamore/Collection-Server/src/database/models"
)

// CharCtrl : Characters controller
type CharCtrl struct{}

type charsRequest struct {
	Characters []models.Character `json:"characters"`
}

// GetUsersCharacters : GetUsersCharacters
// @Summary 특정 유저가 가진 모든 캐릭터들을 가져온다
// @Param memberId path string true "멤버 아이디"
// @Success 0 {object} response.JSONResult{data=[]models.Character}
// @Failure 9011
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
// @Param body body charsRequest true "캐릭터 모델"
// @Success 0
// @Failure 9011
// @Router /members/{memberId}/characters [post]
func (CharCtrl) CreateUsersCharacters(c echo.Context) error {
	var res response.JSONResult
	var req charsRequest
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.Init(c).ValidMember(memberID); err != nil {
		return err
	}

	c.Bind(&req)
	for _, char := range req.Characters {
		char.MemberID = memberID
		if err := db.Create(&char).Error; err != nil {
			if notExist := db.NewRecord(char); notExist {
				code, message := database.ParseError(err)
				res.Code = code
				res.Message = message
				return c.JSON(http.StatusOK, res)
			}
		}
	}
	res.Message = response.StatusText(response.StatusCreated)

	return c.JSON(http.StatusOK, res)
}

// UpdateUsersCharacters : UpdateUsersCharacters
// @Summary 특정 유저가 가진 캐릭터들의 정보를 수정한다
// @Param memberId path string true "멤버 아이디"
// @Param body body charsRequest true "캐릭터 모델"
// @Success 0
// @Failure 9001
// @Failure 9011
// @Router /members/{memberId}/characters [put]
func (CharCtrl) UpdateUsersCharacters(c echo.Context) error {
	var res response.JSONResult
	var req charsRequest
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.Init(c).ValidMember(memberID); err != nil {
		return err
	}

	c.Bind(&req)
	for _, after := range req.Characters {
		after.MemberID = memberID
		var before models.Character
		// 존재하는지 체크
		if err := db.Where("id = ?", after.ID).First(&before).Error; gorm.IsRecordNotFoundError(err) {
			res.Code = response.StatusNotExist
			res.Message = response.StatusText(res.Code)
			return c.JSON(http.StatusOK, res)
		}
		// 업데이트 시도
		if err := db.Model(&before).Update(after).Error; err != nil {
			code, message := database.ParseError(err)
			res.Code = code
			res.Message = message
			return c.JSON(http.StatusOK, res)
		}
	}
	res.Message = response.StatusText(response.StatusUpdated)

	return c.JSON(http.StatusOK, res)
}

// DeleteUsersCharacters : DeleteUsersCharacters
// @Summary 특정 유저가 가진 캐릭터들을 삭제한다
// @Param memberId path string true "멤버 아이디"
// @Param body body charsRequest true "캐릭터 모델"
// @Success 0
// @Failure 9011
// @Router /members/{memberId}/characters [delete]
func (CharCtrl) DeleteUsersCharacters(c echo.Context) error {
	var res response.JSONResult
	var req charsRequest
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.Init(c).ValidMember(memberID); err != nil {
		return err
	}

	c.Bind(&req)
	for _, char := range req.Characters {
		char.MemberID = memberID
		if err := db.Delete(char).Error; err != nil {
			code, message := database.ParseError(err)
			res.Code = code
			res.Message = message
			return c.JSON(http.StatusOK, res)
		}
	}
	res.Message = response.StatusText(response.StatusDeleted)

	return c.JSON(http.StatusOK, res)
}
