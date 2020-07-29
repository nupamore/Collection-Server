package services

import (
	"github.com/jinzhu/gorm"
	"github.com/nupamore/Collection-Server/src/controllers/response"
	"github.com/nupamore/Collection-Server/src/database"
	"github.com/nupamore/Collection-Server/src/database/models"
)

// ValidMember : Check exist member and return info
func ValidMember(memberID string) (models.Member, response.JSONResult) {
	db := database.DB
	var res response.JSONResult
	var member models.Member

	if err := db.Where("id = ?", memberID).First(&member).Error; gorm.IsRecordNotFoundError(err) {
		res.Code = response.StatusNotExistMember
		res.Message = response.StatusText(res.Code)
		return member, res
	}

	return member, res
}
