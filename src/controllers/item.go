package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/controllers/response"
	"github.com/nupamore/Collection-Server/src/database/models"
	"github.com/nupamore/Collection-Server/src/services"
)

// Item : Items controller
type Item struct{}

type itemsRequest struct {
	Items []models.Item `json:"items"`
}

// GetUsersItems : GetUsersItems
// @Summary 특정 유저가 가진 모든 아이템들을 가져온다
// @Param memberId path string true "멤버 아이디"
// @Success 0 {object} response.JSONResult{data=[]models.Item}
// @Failure 9011
// @Router /members/{memberId}/items [get]
func (Item) GetUsersItems(c echo.Context) error {
	var res response.JSONResult
	var items []models.Item
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.ValidMember(memberID); err.Code != 0 {
		return c.JSON(http.StatusOK, err)
	}

	db.Where("memberId = ? AND stackNum > 0", memberID).Find(&items)
	res.Data = items

	return c.JSON(http.StatusOK, res)
}

// AddUsersItems : AddUsersItems
// @Summary 특정 유저가 가진 아이템들의 개수를 추가한다
// @Param memberId path string true "멤버 아이디"
// @Param body body itemsRequest true "아이템 모델"
// @Success 0
// @Failure 9011
// @Router /members/{memberId}/items [post]
func (Item) AddUsersItems(c echo.Context) error {
	var res response.JSONResult
	var req itemsRequest
	var befores, afters []models.Item
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.ValidMember(memberID); err.Code != 0 {
		return c.JSON(http.StatusOK, err)
	}

	c.Bind(&req)
	afters = req.Items
	db.Where("memberId = ?", memberID).Find(&befores)
	for _, after := range afters {
		var before models.Item
		after.MemberID = memberID
		if err := db.Where("itemKey = ?", after.ItemKey).First(&before).Error; gorm.IsRecordNotFoundError(err) {
			// 기존 정보가 없으면 레코드 추가
			db.Create(&after)
		} else {
			// 기존 정보가 있으면 레코드 업데이트
			db.Model(&before).Update(models.Item{
				StackNum: before.StackNum + after.StackNum,
			})
		}
	}
	res.Message = response.StatusText(response.StatusUpdated)

	return c.JSON(http.StatusOK, res)
}

// SubtractUsersItems : SubtractUsersItems
// @Summary 특정 유저가 가진 아이템들의 개수를 뺀다
// @Param memberID path string true "멤버 아이디"
// @Param body body itemsRequest true "아이템 모델"
// @Success 0
// @Failure 9011
// @Router /members/{memberId}/items [delete]]
func (Item) SubtractUsersItems(c echo.Context) error {
	var res response.JSONResult
	var req itemsRequest
	var befores, afters []models.Item
	memberID := c.Param("memberId")
	db, _ := c.Get("db").(*gorm.DB)

	if _, err := services.ValidMember(memberID); err.Code != 0 {
		return c.JSON(http.StatusOK, err)
	}

	c.Bind(&req)
	afters = req.Items
	db.Where("memberId = ?", memberID).Find(&befores)
	for _, after := range afters {
		var before models.Item
		// 없으면 에러
		if err := db.Where("itemKey = ?", after.ItemKey).First(&before).Error; gorm.IsRecordNotFoundError(err) {
			res.Code = response.StatusNotExist
			res.Message = response.StatusText(res.Code)
			return c.JSON(http.StatusOK, res)
		}
		// 있으면 레코드 업데이트
		num := before.StackNum - after.StackNum
		if num < 0 {
			num = 0
		}
		db.Model(&before).Update(models.Item{
			StackNum: num,
		})

	}
	res.Message = response.StatusText(response.StatusUpdated)

	return c.JSON(http.StatusOK, res)
}
