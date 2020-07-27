package models

import "time"

// Character : Character model
type Character struct {
	EvAttack       int       `gorm:"column:EvAttack" json:"EvAttack"`
	EvDefense      int       `gorm:"column:EvDefense" json:"EvDefense"`
	EvHealth       int       `gorm:"column:EvHealth" json:"EvHealth"`
	EvSpAttack     int       `gorm:"column:EvSpAttack" json:"EvSpAttack"`
	EvSpDefense    int       `gorm:"column:EvSpDefense" json:"EvSpDefense"`
	EvSpeed        int       `gorm:"column:EvSpeed" json:"EvSpeed"`
	CharacterExp   int       `gorm:"column:characterExp" json:"characterExp"`
	CharacterID    int       `gorm:"column:characterId" json:"characterId"`
	CharacterKey   int       `gorm:"column:characterKey" json:"characterKey"`
	CharacterLevel int       `gorm:"column:characterLevel" json:"characterLevel"`
	CreatedAt      time.Time `gorm:"column:createdAt" json:"createdAt"`
	ID             int       `gorm:"column:id;primary_key" json:"id"`
	JoinedParty    int       `gorm:"column:joinedParty" json:"joinedParty"`
	JoinedSlot     int       `gorm:"column:joinedSlot" json:"joinedSlot"`
	MemberID       string    `gorm:"column:memberId" json:"memberId"`
	Skill1Level    int       `gorm:"column:skill1Level" json:"skill1Level"`
	Skill2Level    int       `gorm:"column:skill2Level" json:"skill2Level"`
	Skill3Level    int       `gorm:"column:skill3Level" json:"skill3Level"`
	Skill4Level    int       `gorm:"column:skill4Level" json:"skill4Level"`
	UpdatedAt      time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

// TableName sets the insert table name for this struct type
func (c *Character) TableName() string {
	return "haveCharacters"
}
