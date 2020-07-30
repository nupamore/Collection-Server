package services

import (
	"github.com/jinzhu/gorm"
	"github.com/nupamore/Collection-Server/src/controllers/response"
	"github.com/nupamore/Collection-Server/src/database"
	"github.com/nupamore/Collection-Server/src/database/models"
)

var errRes = response.JSONResult{
	Code:    response.StatusNotExistMember,
	Message: response.StatusText(response.StatusNotExistMember),
}

// ValidMember : Check exist member and return info
func ValidMember(memberID string) (member models.Member, res *response.JSONResult) {
	db := database.DB

	if err := db.Where("id = ?", memberID).First(&member).Error; gorm.IsRecordNotFoundError(err) {
		return member, &errRes
	}

	return member, nil
}
