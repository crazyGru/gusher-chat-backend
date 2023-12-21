package block

import "time"

const TableNameUserBlock = "user_block"

type UserBlockDto struct {
	UserID    int       `gorm:"column:user" json:"user"`
	BlockedID int       `gorm:"column:blocked_user" json:"blocked_user"`
	Date      time.Time `gorm:"column:date" json:"date"`
	Note      string    `gorm:"column:note" json:"note"`
}

func (*UserBlockDto) TableName() string {
	return TableNameUserBlock
}
