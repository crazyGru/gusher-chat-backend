package repo

import (
	"backend/protocol"
	model "backend/protocol/protocol"
	"fmt"
	"gorm.io/gorm"
)

type ProtocolRepository struct {
	db *gorm.DB
}

func NewProtocolRepository(db *gorm.DB) *ProtocolRepository {
	return &ProtocolRepository{
		db: db,
	}
}

func (m *ProtocolRepository) query(params protocol.ProtocolSearch) *gorm.DB {
	q := m.db.Model(model.ProtocolDto{})
	if params.Id != 0 {
		q.Where("id = ?", params.Id)
	}
	if params.Name != "" {
		q.Where("name = ?", params.Name)
	}
	return q
}

func (m *ProtocolRepository) Find(params protocol.ProtocolSearch) ([]*model.Protocol, error) {
	var result []*model.Protocol

	q := m.query(params)
	var rows []*model.ProtocolDto
	q.Preload("Tasks").Find(&rows)
	if q.Error != nil {
		return nil, q.Error
	}
	for i := range rows {
		result = append(result, model.FromDto(rows[i]))
	}
	return result, nil
}

func (m *ProtocolRepository) Count(params protocol.ProtocolSearch) (int, error) {
	var cnt int64
	q := m.query(params)
	q.Count(&cnt)
	if q.Error != nil {
		return 0, q.Error
	}
	return int(cnt), nil
}

func (m *ProtocolRepository) Delete(item *model.Protocol) {
	dto := model.ToDto(item)
	m.db.Delete(dto)
}

func (m *ProtocolRepository) err(e error) error {
	return fmt.Errorf("protocol.ProtocolRepository: %w", e)
}

func (m *ProtocolRepository) Save(item *model.Protocol) error {
	dto := model.ToDto(item)
	if dto.ID != 0 {
		m.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(dto)
	} else {
		m.db.Create(dto)
	}
	*item = *model.FromDto(dto)
	return nil

	//insertProtocolQuery := "insert into protocol (name) values (?)"
	//updateProtocolQuery := "update protocol set name = ? where id = ?"
	//var protocolId int
	//
	//if protocol.id == 0 {
	//	result, err := m.db.Exec(insertProtocolQuery, protocol.Name)
	//	if err != nil {
	//		return err
	//	}
	//	if newId, err := result.LastInsertId(); err != nil {
	//		return err
	//	} else {
	//		protocolId = int(newId)
	//	}
	//} else {
	//	protocolId = protocol.id
	//	if _, err := m.db.Exec(updateProtocolQuery, protocol.Name, protocol.id); err != nil {
	//		return err
	//	}
	//}
	//if err := m.saveTasks(protocolId, protocol.Tasks); err != nil {
	//	return err
	//}
	//return nil
}

func (m *ProtocolRepository) saveTasks(protocolId int, tasks []*model.Task) error {
	panic("implement me")
	//
	//insertTaskQuery := "insert into protocol_task " +
	//	"(protocol_id, role, complete_fn, depends_on, title, content) " +
	//	"values (:protocol_id, :role, :complete_fn, :depends_on, :title, :content)"
	//updateTaskQuery := "update protocol_task " +
	//	"set protocol_id = :protocol_id, role = :role, complete_fn = :complete_fn, " +
	//	"depends_on = :depends_on, title = :title, content = :content" +
	//	"where id = :id"
	//
	//for _, task := range tasks {
	//	data := map[string]any{
	//		"protocol_id": protocolId,
	//		"role":        string(task.StartupRole),
	//		"complete_fn": task.CompleteFn,
	//		"title":       task.Title,
	//		"content":     task.Content,
	//	}
	//	if dependsOn, err := json.Marshal(task.DependsOn); err == nil {
	//		data["depends_on"] = string(dependsOn)
	//	} else {
	//		return err
	//	}
	//	if task.id == 0 {
	//		if _, err := m.db.NamedExec(insertTaskQuery, data); err != nil {
	//			return err
	//		}
	//	} else {
	//		data["id"] = task.id
	//		if _, err := m.db.NamedExec(updateTaskQuery, data); err != nil {
	//			return err
	//		}
	//	}
	//}
	//return nil
}

func (m *ProtocolRepository) FindStartupProtocol(startupId int, protocolId int) (*protocol.StartupProtocol, error) {
	panic("implement me")
	//var completeTasks []int
	//rows, err := m.db.Queryx("select * from protocol_task_startup where startup_id = ?", startupId)
	//if err != nil {
	//	return nil, err
	//}
	//var taskId int
	//for rows.Next() {
	//	err := rows.Scan(&taskId)
	//	if err != nil {
	//		return nil, m.err(err)
	//	}
	//	completeTasks = append(completeTasks, taskId)
	//}
	//p, err := m.Find(protocol.ProtocolSearch{id: protocolId})
	//if err != nil {
	//	return nil, m.err(err)
	//}
	//if len(p) != 1 {
	//	return nil, m.err(fmt.Errorf("failed to load protocol id=%d", protocolId))
	//}
	//sp := &protocol.StartupProtocol{
	//	StartupId:     startupId,
	//	Protocol:      p[0],
	//	CompleteTaskIds: completeTasks,
	//}
	//return sp, nil
	//
}

func (m *ProtocolRepository) SaveStartupProtocol(startupProtocol *protocol.StartupProtocol) error {
	panic("implement me")
	//var err error
	//tx, err := m.db.Beginx()
	//if err != nil {
	//	return m.err(err)
	//}
	//_, err = tx.Exec("delete from protocol_task_startup where startup_id = ?", startupProtocol.StartupId)
	//if err != nil {
	//	tx.Rollback()
	//	return m.err(err)
	//}
	//for _, taskId := range startupProtocol.CompleteTaskIds {
	//	_, err := tx.Exec("insert into protocol_task_startup (startup_id, task_id) values (?,?)", startupProtocol.StartupId, taskId)
	//	if err != nil {
	//		tx.Rollback()
	//		return err
	//	}
	//}
	//err = tx.Commit()
	//if err != nil {
	//	return err
	//}
	//return nil

}
