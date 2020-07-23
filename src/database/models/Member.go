package models

import (
	"time"
)

// Member : Member model
type Member struct {
	BattleClearCount     int       `gorm:"column:battleClearCount" json:"battleClearCount"`
	BattleSpeed          int       `gorm:"column:battleSpeed" json:"battleSpeed"`
	Berry                int       `gorm:"column:berry" json:"berry"`
	CreatedAt            time.Time `gorm:"column:createdAt" json:"createdAt"`
	Exp                  int       `gorm:"column:exp" json:"exp"`
	GetCharacterCount    int       `gorm:"column:getCharacterCount" json:"getCharacterCount"`
	ID                   string    `gorm:"column:id;primary_key" json:"id"`
	Index                string    `gorm:"column:index" json:"index"`
	LastSelectedChapter  int       `gorm:"column:lastSelectedChapter" json:"lastSelectedChapter"`
	LastSelectedParty    int       `gorm:"column:lastSelectedParty" json:"lastSelectedParty"`
	Level                int       `gorm:"column:level" json:"level"`
	MaxClearBattleStage  int       `gorm:"column:maxClearBattleStage" json:"maxClearBattleStage"`
	MaxHaveCharactersNum int       `gorm:"column:maxHaveCharactersNum" json:"maxHaveCharactersNum"`
	MaxHaveEquipmentsNum int       `gorm:"column:maxHaveEquipmentsNum" json:"maxHaveEquipmentsNum"`
	MaxOpenedChapter     int       `gorm:"column:maxOpenedChapter" json:"maxOpenedChapter"`
	Money                int       `gorm:"column:money" json:"money"`
	Nickname             string    `gorm:"column:nickname" json:"nickname"`
	Stardust             int       `gorm:"column:stardust" json:"stardust"`
	UpdatedAt            time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

// TableName sets the insert table name for this struct type
func (m *Member) TableName() string {
	return "members"
}
