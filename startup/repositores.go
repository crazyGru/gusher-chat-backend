package startup

import (
	"backend/shared"
	"backend/startup/applicant"
	"backend/startup/role"
	"backend/startup/startup"
)

type StartupRepository interface {
	UserRolesInStartup(startupId int, userId int) []shared.StartupRole
	Find(params Search) ([]*startup.Startup, error)
	Count(params Search) (int, error)
	Delete(item *startup.Startup) error
	FindApplicants(params ApplicantSearch) ([]*applicant.Applicant, error)
	SaveStartup(item *startup.Startup)
}

type RoleRepository interface {
	Find(params RoleSearch) ([]*role.Role, error)
	Count(params RoleSearch) (int, error)
	Save(r *role.Role)
	Delete(a *applicant.Applicant)
}

type ApplicantRepository interface {
	Find(params ApplicantSearch) ([]*applicant.Applicant, error)
	Count(params ApplicantSearch) (int, error)
	Save(a *applicant.Applicant)
	Delete(a *applicant.Applicant) error
}
