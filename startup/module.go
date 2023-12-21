package startup

import (
	"backend/shared"
	"backend/startup/applicant"
	"backend/startup/role"
	"backend/startup/startup"
	"errors"
)

type Module struct {
	startupRepository   StartupRepository
	roleRepository      RoleRepository
	applicantRepository ApplicantRepository
}

func NewModule(repo StartupRepository, roleRepo RoleRepository, applicantRepository ApplicantRepository) *Module {
	return &Module{
		startupRepository:   repo,
		roleRepository:      roleRepo,
		applicantRepository: applicantRepository,
	}
}

func (m *Module) StartupProtocolId(startupId int) (int, error) {
	items, err := m.startupRepository.Find(Search{Id: startupId})
	if err != nil {
		return 0, err
	}
	if len(items) == 0 {
		return 0, errors.New("startup not found")
	}
	return items[0].Protocol(), nil
}

func (m *Module) FindStartups(params Search) ([]*startup.Startup, error) {
	return m.startupRepository.Find(params)
}

func (m *Module) CountStartups(params Search) (int, error) {
	return m.startupRepository.Count(params)
}

func (m *Module) SaveStartup(item *startup.Startup) {
	m.startupRepository.SaveStartup(item)
}

func (m *Module) FindRoles(params RoleSearch) ([]*role.Role, error) {
	return m.roleRepository.Find(params)
}

func (m *Module) FindRole(params RoleSearch) (*role.Role, error) {
	r, err := m.FindRoles(params)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, errors.New("not found")
	}
	return r[0], nil
}

func (m *Module) SaveRole(r *role.Role) {
	m.roleRepository.Save(r)
}

func (m *Module) FindApplicants(params ApplicantSearch) ([]*applicant.Applicant, error) {
	return m.applicantRepository.Find(params)
}

// todo need error handling here
func (m *Module) UserRolesInStartup(startupId int, userId int) []shared.StartupRole {
	var result []shared.StartupRole
	var err error
	var s []*startup.Startup
	s, err = m.FindStartups(Search{Id: startupId})
	if err != nil || len(s) == 0 {
		return result
	}
	if s[0].User() == userId {
		result = append(result, "founder")
	}
	var roles []*role.Role
	roles, err = m.FindRoles(RoleSearch{Startup: startupId})
	if err != nil || len(roles) == 0 {
		return result
	}
	var applicants []*applicant.Applicant
	applicants, err = m.FindApplicants(ApplicantSearch{User: userId, Startup: startupId, Status: "joined"})
	for _, a := range applicants {
		if r, err := m.FindRole(RoleSearch{Id: a.Role}); err == nil {
			result = append(result, r.ProtocolRole()...)
		}
	}
	return result
}
