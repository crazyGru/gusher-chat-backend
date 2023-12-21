package repo

import (
	"gorm.io/gorm"
)

type ProtocolTaskStartupRepository struct {
	db *gorm.DB
}

func NewProtocolTaskStartupRepository(db *gorm.DB) *ProtocolTaskStartupRepository {
	return &ProtocolTaskStartupRepository{db: db}
}

func (p *ProtocolTaskStartupRepository) Find(startupId int) ([]int, error) {
	var result []int
	tx := p.db.Raw("select task_id from protocol_task_startup where startup_id = ?", startupId).Scan(&result)
	if tx.Error != nil {
		return result, tx.Error
	}
	return result, nil
}

func (p *ProtocolTaskStartupRepository) Save(startupId int, tasks []int) {
	p.db.Exec("delete from protocol_task_startup where startup_id = ?", startupId)
	for _, taskId := range tasks {
		p.db.Exec("insert into protocol_task_startup (startup_id, task_id) values (?,?)", startupId, taskId)
	}
}
