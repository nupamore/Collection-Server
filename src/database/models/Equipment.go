package models

import "time"

// Equipment : Equipment model
type Equipment struct {
	CreatedAt        time.Time `gorm:"column:createdAt" json:"createdAt"`
	EquipmentID      int       `gorm:"column:equipmentId" json:"equipmentId"`
	EquipmentKey     int       `gorm:"column:equipmentKey" json:"equipmentKey"`
	EquipmentLevel   int       `gorm:"column:equipmentLevel" json:"equipmentLevel"`
	ID               int       `gorm:"column:id;primary_key" json:"id"`
	MemberID         string    `gorm:"column:memberId" json:"memberId"`
	OwnerCharacterID int       `gorm:"column:ownerCharacterId" json:"ownerCharacterId"`
	UpdatedAt        time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

// TableName sets the insert table name for this struct type
func (e *Equipment) TableName() string {
	return "haveEquipments"
}
