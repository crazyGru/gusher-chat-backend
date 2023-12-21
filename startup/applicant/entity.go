package applicant

type Applicant struct {
	Id   int `db:"id"`
	Role int `db:"role"`
	User int `db:"user"`
}

func FromDto(dto *ApplicantDto) *Applicant {
	return &Applicant{
		Id:   *dto.ID,
		Role: *dto.Role,
		User: *dto.User,
	}
}

func ToDto(a *Applicant) *ApplicantDto {
	return &ApplicantDto{
		ID:   &a.Id,
		Role: &a.Role,
		User: &a.User,
	}
}
