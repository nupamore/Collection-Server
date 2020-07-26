package models

import "time"

// Item : Item model
type Item struct {
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	ItemKey   int       `gorm:"column:itemKey" json:"itemKey"`
	MemberID  string    `gorm:"column:memberId" json:"memberId"`
	StackNum  int       `gorm:"column:stackNum" json:"stackNum"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

// TableName sets the insert table name for this struct type
func (i *Item) TableName() string {
	return "haveItems"
}
