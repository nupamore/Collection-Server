package models

// Quest : Quest model
type Quest struct {
	Completed  int64  `gorm:"column:completed" json:"completed"`
	CurrentNum int64  `gorm:"column:currentNum" json:"currentNum"`
	ID         int    `gorm:"column:id;primary_key" json:"id"`
	MemberID   string `gorm:"column:memberId" json:"memberId"`
	QuestKey   int64  `gorm:"column:questKey" json:"questKey"`
}

// TableName sets the insert table name for this struct type
func (q *Quest) TableName() string {
	return "quests"
}
