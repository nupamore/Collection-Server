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

// Quest : Quests controller
type Quest struct{}

type questsRequest struct {
	Quests []models.Quest `json:"quests"`
}

// GetUsersQuests : GetUsersQuests
// @Summary 특정 유저가 가진 모든 퀘스트들을 가져온다
// @Param memberId path string true "멤버 아이디"
// @Success 0 {object} response.JSONResult{data=[]models.Quest}
// @Failure 9011
// @Router /members/{memberId}/quests [get]
func (Quest) GetUsersQuests(c echo.Context) error {
	var res response.JSONResult
	var quests []models.Quest
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.ValidMember(memberID); err.Code != 0 {
		return c.JSON(http.StatusOK, err)
	}

	db.Where("memberId = ?", memberID).Find(&quests)
	res.Data = quests

	return c.JSON(http.StatusOK, res)
}

// CreateUsersQuests : CreateUsersQuests
// @Summary 특정 유저에게 퀘스트들을 할당해준다
// @Param memberId path string true "멤버 아이디"
// @Param body body questsRequest true "퀘스트 모델"
// @Success 0
// @Failure 9011
// @Router /members/{memberId}/quests [post]
func (Quest) CreateUsersQuests(c echo.Context) error {
	var res response.JSONResult
	var req questsRequest
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.ValidMember(memberID); err.Code != 0 {
		return c.JSON(http.StatusOK, err)
	}

	c.Bind(&req)
	for _, quest := range req.Quests {
		quest.MemberID = memberID
		if err := db.Create(&quest).Error; err != nil {
			if notExist := db.NewRecord(quest); notExist {
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

// UpdateUsersQuests : UpdateUsersQuests
// @Summary 특정 유저가 가진 퀘스트들의 정보를 수정한다
// @Param memberId path string true "멤버 아이디"
// @Param body body questsRequest true "퀘스트 모델"
// @Success 0
// @Failure 9001
// @Failure 9011
// @Router /members/{memberId}/quests [put]
func (Quest) UpdateUsersQuests(c echo.Context) error {
	var res response.JSONResult
	var req questsRequest
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.ValidMember(memberID); err.Code != 0 {
		return c.JSON(http.StatusOK, err)
	}

	c.Bind(&req)
	for _, after := range req.Quests {
		after.MemberID = memberID
		var before models.Quest
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

// DeleteUsersQuests : DeleteUsersQuests
// @Summary 특정 유저가 가진 퀘스트들을 삭제한다
// @Param memberId path string true "멤버 아이디"
// @Param body body questsRequest true "퀘스트 모델"
// @Success 0
// @Failure 9011
// @Router /members/{memberId}/quests [delete]
func (Quest) DeleteUsersQuests(c echo.Context) error {
	var res response.JSONResult
	var req questsRequest
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.ValidMember(memberID); err.Code != 0 {
		return c.JSON(http.StatusOK, err)
	}

	c.Bind(&req)
	for _, Quest := range req.Quests {
		Quest.MemberID = memberID
		if err := db.Delete(Quest).Error; err != nil {
			code, message := database.ParseError(err)
			res.Code = code
			res.Message = message
			return c.JSON(http.StatusOK, res)
		}
	}
	res.Message = response.StatusText(response.StatusDeleted)

	return c.JSON(http.StatusOK, res)
}
