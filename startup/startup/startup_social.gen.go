// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package startup

import (
	"time"
)

const TableNameStartupSocial = "startup_social"

// StartupSocial mapped from table <startup_social>
type StartupSocialDto struct {
	Startup       int       `gorm:"column:startup" json:"startup"`
	SocialNetwork int       `gorm:"column:social_network" json:"social_network"`
	Reference     string    `gorm:"column:reference" json:"reference"`
	Date          time.Time `gorm:"column:date" json:"date"`
	Verified      int       `gorm:"column:verified" json:"verified"`
}

// TableName StartupSocial's table name
func (*StartupSocialDto) TableName() string {
	return TableNameStartupSocial
}
