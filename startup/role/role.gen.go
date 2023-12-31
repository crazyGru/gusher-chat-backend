// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package role

import (
	"time"
)

const TableNameRole = "role"

// Role mapped from table <role>
type RoleDto struct {
	ID                int        `gorm:"column:id" json:"id"`
	Startup           int        `gorm:"column:startup" json:"startup"`
	Title             string     `gorm:"column:title" json:"title"`
	Description       string     `gorm:"column:description" json:"description"`
	Equity            float64    `gorm:"column:equity" json:"equity"`
	Applicants        int        `gorm:"column:applicants" json:"applicants"`
	SelectedApplicant int        `gorm:"column:selected_applicant" json:"selected_applicant"`
	Views             int        `gorm:"column:views" json:"views"`
	Status            string     `gorm:"column:status" json:"status"`
	Vanity            string     `gorm:"column:vanity" json:"vanity"`
	Created           time.Time  `gorm:"column:created" json:"created"`
	Modified          time.Time  `gorm:"column:modified" json:"modified"`
	Opened            time.Time  `gorm:"column:opened" json:"opened"`
	Closed            time.Time  `gorm:"column:closed" json:"closed"`
	Filled            time.Time  `gorm:"column:filled" json:"filled"`
	Awarded           time.Time  `gorm:"column:awarded" json:"awarded"`
	Blocked           time.Time  `gorm:"column:blocked" json:"blocked"`
	Hidden            time.Time  `gorm:"column:hidden" json:"hidden"`
	Featured          int        `gorm:"column:featured" json:"featured"`
	ProtocolRole      string     `gorm:"column:protocol_role;not null" json:"protocol_role"`
	Goals             []*GoalDto `gorm:"foreignkey:role"`
}

// TableName Role's table name
func (*RoleDto) TableName() string {
	return TableNameRole
}
