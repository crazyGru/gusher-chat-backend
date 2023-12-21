package repo

import (
	"backend/startup"
	"backend/startup/applicant"
	"gorm.io/gorm"
)

type ApplicantRepository struct {
	db *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) *ApplicantRepository {
	return &ApplicantRepository{db: db}
}

func (r *ApplicantRepository) Delete(a *applicant.Applicant) error {
	//TODO implement me
	panic("implement me")
}

func (r *ApplicantRepository) query(params startup.ApplicantSearch) *gorm.DB {
	q := r.db.Model(applicant.ApplicantDto{})
	if params.Id != 0 {
		q.Where("applicant.id = ?", params.Id)
	}
	if params.User != 0 {
		q.Where("applicant.user = ?", params.User)
	}
	if params.Role != 0 {
		q.Where("applicant.role = ?", params.Role)
	}
	if params.Status != "" {
		q.Where("applicant.status = ?", params.Status)
	}
	if params.Startup != 0 {
		q.Joins("left join role on role.id = applicant.role")
		q.Where("role.startup = ?", params.Startup)
	}
	return q
}

func (r *ApplicantRepository) Find(params startup.ApplicantSearch) ([]*applicant.Applicant, error) {
	var rows []*applicant.ApplicantDto
	q := r.query(params)
	q.Find(&rows)
	if q.Error != nil {
		return nil, q.Error
	}
	var result []*applicant.Applicant
	for _, row := range rows {
		result = append(result, applicant.FromDto(row))
	}
	return result, nil
}

func (r *ApplicantRepository) Count(params startup.ApplicantSearch) (int, error) {
	var cnt int64
	q := r.query(params)
	q.Count(&cnt)
	if q.Error != nil {
		return 0, q.Error
	}
	return int(cnt), nil
}

func (r *ApplicantRepository) Save(a *applicant.Applicant) {
	//TODO implement me
	panic("implement me")
}
