// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package protocol

const TableNameProtocolTask = "protocol_task"

// ProtocolTask mapped from table <protocol_task>
type TaskDto struct {
	ID         int     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProtocolID int     `gorm:"column:protocol_id;not null" json:"protocol_id"`
	Role       string  `gorm:"column:role" json:"role"`
	CompleteFn string  `gorm:"column:complete_fn;not null" json:"complete_fn"`
	DependsOn  DepList `gorm:"column:depends_on;serializer:json" json:"depends_on"`
	Title      string  `gorm:"column:title" json:"title"`
	Content    string  `gorm:"column:content" json:"content"`
}

// TableName ProtocolTask's table name
func (*TaskDto) TableName() string {
	return TableNameProtocolTask
}
