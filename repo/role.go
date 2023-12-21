package repo

import (
	"backend/shared"
	"backend/startup"
	"backend/startup/applicant"
	"backend/startup/role"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (s *RoleRepository) Delete(a *applicant.Applicant) {
	//TODO implement me
	panic("implement me")
}

func (s *RoleRepository) query(params startup.RoleSearch) *gorm.DB {
	q := s.db.Preload("Goals").Model(role.RoleDto{})
	if params.Id != 0 {
		q.Where("id = ?", params.Id)
	}
	if params.Startup != 0 {
		q.Where("startup = ?", params.Startup)
	}
	return q
}

func (s *RoleRepository) Find(params startup.RoleSearch) ([]*role.Role, error) {
	q := s.query(params)
	if params.Offset != 0 {
		q.Offset(params.Offset)
	}
	if params.Limit != 0 {
		q.Limit(params.Limit)
	}
	var rows []*role.RoleDto
	q.Find(&rows)
	if q.Error != nil {
		return nil, q.Error
	}
	var result []*role.Role
	for i := range rows {

		result = append(result, role.FromDto(rows[i]))
	}
	return result, nil
}

func (s *RoleRepository) Count(params startup.RoleSearch) (int, error) {
	q := s.query(params)
	var cnt int64
	q.Count(&cnt)
	if q.Error != nil {
		return 0, q.Error
	}
	return int(cnt), nil
}

func (s *RoleRepository) Save(r *role.Role) {
	roleDto := role.ToDto(r)
	tx := s.db.Save(roleDto)
	if tx.Error != nil {
		shared.ErrorLogger.Println(tx.Error.Error())
	}
}
