package models

import (
	"database/sql"
	"time"
)

// Member : Member model
type Member struct {
	BattleClearCount     int            `gorm:"column:battleClearCount"`
	BattleSpeed          int            `gorm:"column:battleSpeed"`
	Berry                int            `gorm:"column:berry"`
	CreatedAt            time.Time      `gorm:"column:createdAt"`
	Exp                  int            `gorm:"column:exp"`
	GetCharacterCount    int            `gorm:"column:getCharacterCount"`
	ID                   string         `gorm:"column:id;primary_key"`
	Index                sql.NullString `gorm:"column:index"`
	LastSelectedChapter  int            `gorm:"column:lastSelectedChapter"`
	LastSelectedParty    int            `gorm:"column:lastSelectedParty"`
	Level                int            `gorm:"column:level"`
	MaxClearBattleStage  int            `gorm:"column:maxClearBattleStage"`
	MaxHaveCharactersNum int            `gorm:"column:maxHaveCharactersNum"`
	MaxHaveEquipmentsNum int            `gorm:"column:maxHaveEquipmentsNum"`
	MaxOpenedChapter     int            `gorm:"column:maxOpenedChapter"`
	Money                int            `gorm:"column:money"`
	Nickname             sql.NullString `gorm:"column:nickname"`
	Stardust             int            `gorm:"column:stardust"`
	UpdatedAt            time.Time      `gorm:"column:updatedAt"`
}
