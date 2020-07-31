package controllers

import (
	"net/http"

	"github.com/fatih/structs"
	"github.com/nupamore/Collection-Server/src/database"
	"github.com/nupamore/Collection-Server/src/services"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/controllers/response"
	"github.com/nupamore/Collection-Server/src/database/models"
)

// Equip : Equipments controller
type Equip struct{}

type equipsRequest struct {
	Equipments []models.Equipment `json:"equipments"`
}

// GetUsersEquipments : GetUsersEquipments
// @Summary 특정 유저가 가진 모든 장비들을 가져온다
// @Param memberId path string true "멤버 아이디"
// @Success 0 {object} response.JSONResult{data=[]models.Equipment}
// @Failure 9011
// @Router /members/{memberId}/equipments [get]
func (Equip) GetUsersEquipments(c echo.Context) error {
	var res response.JSONResult
	var equips []models.Equipment
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.ValidMember(memberID); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	db.Where("memberId = ?", memberID).Find(&equips)
	res.Data = equips

	return c.JSON(http.StatusOK, res)
}

// CreateUsersEquipments : CreateUsersEquipments
// @Summary 특정 유저에게 장비들을 할당해준다
// @Param memberId path string true "멤버 아이디"
// @Param body body equipsRequest true "장비 모델"
// @Success 0
// @Failure 9011
// @Router /members/{memberId}/equipments [post]
func (Equip) CreateUsersEquipments(c echo.Context) error {
	var res response.JSONResult
	var req equipsRequest
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.ValidMember(memberID); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	c.Bind(&req)
	for _, equip := range req.Equipments {
		equip.MemberID = memberID
		if err := db.Create(&equip).Error; err != nil {
			if notExist := db.NewRecord(equip); notExist {
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

// UpdateUsersEquipments : UpdateUsersEquipments
// @Summary 특정 유저가 가진 장비들의 정보를 수정한다
// @Param memberId path string true "멤버 아이디"
// @Param body body equipsRequest true "장비 모델"
// @Success 0
// @Failure 9001
// @Failure 9011
// @Router /members/{memberId}/equipments [put]
func (Equip) UpdateUsersEquipments(c echo.Context) error {
	var res response.JSONResult
	var req equipsRequest
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.ValidMember(memberID); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	c.Bind(&req)
	for _, after := range req.Equipments {
		after.MemberID = memberID
		var before models.Equipment
		// 존재하는지 체크
		if err := db.Where("id = ?", after.ID).First(&before).Error; gorm.IsRecordNotFoundError(err) {
			res.Code = response.StatusNotExist
			res.Message = response.StatusText(res.Code)
			return c.JSON(http.StatusOK, res)
		}
		// 업데이트 시도
		if err := db.Model(&before).Update(structs.Map(after)).Error; err != nil {
			code, message := database.ParseError(err)
			res.Code = code
			res.Message = message
			return c.JSON(http.StatusOK, res)
		}
	}
	res.Message = response.StatusText(response.StatusUpdated)

	return c.JSON(http.StatusOK, res)
}

// DeleteUsersEquipments : DeleteUsersEquipments
// @Summary 특정 유저가 가진 장비들을 삭제한다
// @Param memberId path string true "멤버 아이디"
// @Param body body equipsRequest true "장비 모델"
// @Success 0
// @Failure 9011
// @Router /members/{memberId}/equipments [delete]
func (Equip) DeleteUsersEquipments(c echo.Context) error {
	var res response.JSONResult
	var req equipsRequest
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.ValidMember(memberID); err != nil {
		return c.JSON(http.StatusOK, err)
	}

	c.Bind(&req)
	for _, equip := range req.Equipments {
		equip.MemberID = memberID
		if err := db.Delete(equip).Error; err != nil {
			code, message := database.ParseError(err)
			res.Code = code
			res.Message = message
			return c.JSON(http.StatusOK, res)
		}
	}
	res.Message = response.StatusText(response.StatusDeleted)

	return c.JSON(http.StatusOK, res)
}
