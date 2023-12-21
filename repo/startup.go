package repo

import (
	"backend/shared"
	startup2 "backend/startup"
	"backend/startup/applicant"
	"backend/startup/role"
	"backend/startup/startup"
	"gorm.io/gorm"
)

type StartupRepository struct {
	db *gorm.DB
}

func (m *StartupRepository) Delete(item *startup.Startup) error {
	//TODO implement me
	panic("implement me")
}

func NewStartupRepository(db *gorm.DB) *StartupRepository {
	return &StartupRepository{db: db}
}

func (m *StartupRepository) query(params startup2.Search) *gorm.DB {
	query := m.db.Model(startup.StartupDto{}).Preload("Topics")
	if params.Id != 0 {
		query.Where("id = ?", params.Id)
	}
	if params.Vanity != "" {
		query.Where("vanity = ?", params.Vanity)
	}
	return query
}

func (m *StartupRepository) Find(params startup2.Search) ([]*startup.Startup, error) {
	var result []*startup.Startup
	var rows []*startup.StartupDto
	query := m.query(params)
	if params.Offset != 0 {
		query.Offset(params.Offset)
	}
	if params.Limit != 0 {
		query.Limit(params.Limit)
	}

	query.Find(&rows)
	for i := range rows {
		result = append(result, startup.FromDto(rows[i]))
	}
	return result, nil
}

func (m *StartupRepository) Count(params startup2.Search) (int, error) {
	query := m.query(params)
	var cnt int64
	query.Count(&cnt)
	if query.Error != nil {
		return 0, query.Error
	}
	return int(cnt), nil
}

func (m *StartupRepository) startupRoles(startupId int) ([]*role.Role, error) {
	var roles []*role.RoleDto
	var result []*role.Role
	m.db.Model(role.RoleDto{}).Where("startup = ?", startupId).Find(&roles)
	for i := range roles {
		result = append(result, role.FromDto(roles[i]))
	}
	return result, nil
}

func (m *StartupRepository) FindApplicants(params startup2.ApplicantSearch) ([]*applicant.Applicant, error) {
	var rows []*applicant.ApplicantDto
	var result []*applicant.Applicant
	q := m.db.Model(applicant.ApplicantDto{})
	if params.Id != 0 {
		q.Where("id = ?", params.Id)
	}
	if params.User != 0 {
		q.Where("user = ?", params.User)
	}
	if params.Role != 0 {
		q.Where("role = ?", params.Role)
	}
	q.Find(&rows)
	for i := range rows {
		result = append(result, applicant.FromDto(rows[i]))
	}
	return result, nil
}

func (m *StartupRepository) UserRolesInStartup(startupId int, userId int) []shared.StartupRole {
	//TODO implement me
	panic("implement me")
}

func (m *StartupRepository) SaveStartup(item *startup.Startup) {
	startupDto := startup.ToDto(item)
	tx := m.db.Save(startupDto)
	if tx.Error != nil {
		shared.ErrorLogger.Println(tx.Error.Error())
	}
}
