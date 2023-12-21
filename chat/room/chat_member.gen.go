package room

import "time"

const TableNameRoomMember = "channel_member"

type RoomMemberDto struct {
	RoomID    int       `gorm:"column:channel" json:"channel"`
	UserID    int       `gorm:"column:user" json "user"`
	Status    string    `gorm:"column:status" json:"status"`
	Activated time.Time `gorm:"column:activated" json:"activated"`
	Blocked   time.Time `gorm:"column:blocked" json:"blocked"`
	IsAdmin   int       `gorm:"column:admin" json:"admin"`
	Notified  int       `gorm:"column:notified" json:"notified"`
}

func (*RoomMemberDto) TableName() string {
	return TableNameRoomMember
}
