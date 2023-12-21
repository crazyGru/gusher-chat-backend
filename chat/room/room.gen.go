package room

import "time"

const TableNameRoom = "channel"

type RoomDto struct {
	Id           int       `gorm:"column:id" json:"id"`
	Startup      int       `gorm:"column:startup" json:startup`
	Name         string    `gorm:"column:name" json:"name"`
	Privacy      string    `gorm:"column:privacy" json:"privacy"`
	Direct       int       `gorm:"column:direct" json:"direct"`
	Admins       int       `gorm:"column:admins" json:"admins"`
	Recruiters   int       `gorm:"column:recruiters" json:"recruiters"`
	Broadcasters int       `gorm:"column:broadcasters" json:"broadcasters"`
	Status       string    `gorm:"column:status" json:"status"`
	Created      time.Time `gorm:"column:created" json:"created"`
	Archived     time.Time `gorm:"column:archived" json:"archived"`
	Hidden       time.Time `gorm:"column:hidden" json:"hidden"`
	Blocked      time.Time `gorm:"column:blocked" json:"blocked"`
}

func (*RoomDto) TableName() string {
	return TableNameRoom
}
