package services

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/nupamore/Collection-Server/src/controllers/response"
	"github.com/nupamore/Collection-Server/src/database/models"
)

// Service : Serivce struct
type Service struct{}

var service Service
var ctx echo.Context

// Init : Init
func Init(c echo.Context) Service {
	ctx = c
	return service
}

// ValidMember : Check exist member and return info
func (Service) ValidMember(memberID string) (models.Member, error) {
	db := ctx.Get("db").(*gorm.DB)
	var res response.JSONResult
	var member models.Member

	if err := db.Where("id = ?", memberID).First(&member).Error; gorm.IsRecordNotFoundError(err) {
		res.Code = response.StatusNotExistMember
		res.Message = response.StatusText(res.Code)
		ctx.JSON(http.StatusOK, res)
		return member, errors.New(res.Message)
	}

	return member, nil
}
