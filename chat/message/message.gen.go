package message

import "time"

const TableNameMessage = "post"

type MessageDto struct {
	ID       int       `gorm:"column:id" json:"id"`
	SenderID int       `gorm:"column:sender" json:"sender"`
	RoomID   int       `gorm:"column:channel" json:"room"`
	Created  time.Time `gorm:"column:created" json:"created"`
	Attach   Attachment
	Text     string `gorm:"column:text" json:"text"`
}

func (*MessageDto) TableName() string {
	return TableNameMessage
}
