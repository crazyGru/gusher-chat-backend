// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package startup

import (
	"time"
)

const TableNameStartupTopic = "startup_topic"

// StartupTopic mapped from table <startup_topic>
type StartupTopicDto struct {
	Startup   int        `gorm:"column:startup" json:"startup"`
	Topic     int        `gorm:"column:topic" json:"topic"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`
}

// TableName StartupTopic's table name
func (*StartupTopicDto) TableName() string {
	return TableNameStartupTopic
}