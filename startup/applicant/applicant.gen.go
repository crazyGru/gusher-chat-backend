// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package applicant

import (
	"time"
)

const TableNameApplicant = "applicant"

// Applicant mapped from table <applicant>
type ApplicantDto struct {
	ID                      *int       `gorm:"column:id" json:"id"`
	Role                    *int       `gorm:"column:role" json:"role"`
	User                    *int       `gorm:"column:user" json:"user"`
	Score                   *int       `gorm:"column:score" json:"score"`
	Scores                  *int       `gorm:"column:scores" json:"scores"`
	Status                  *string    `gorm:"column:status" json:"status"`
	Applied                 *time.Time `gorm:"column:applied" json:"applied"`
	Modified                *time.Time `gorm:"column:modified" json:"modified"`
	Reviewed                *time.Time `gorm:"column:reviewed" json:"reviewed"`
	EntrepreneurSigned      *time.Time `gorm:"column:entrepreneur_signed" json:"entrepreneur_signed"`
	StartupSigned           *time.Time `gorm:"column:startup_signed" json:"startup_signed"`
	Joined                  *time.Time `gorm:"column:joined" json:"joined"`
	AwardUploaded           *time.Time `gorm:"column:award_uploaded" json:"award_uploaded"`
	Withdrawn               *time.Time `gorm:"column:withdrawn" json:"withdrawn"`
	Equity                  *int       `gorm:"column:equity" json:"equity"`
	Notified                *int       `gorm:"column:notified" json:"notified"`
	SignatureRequest        *string    `gorm:"column:signature_request" json:"signature_request"`
	EntrepreneurSignatureID *string    `gorm:"column:entrepreneur_signature_id" json:"entrepreneur_signature_id"`
	EntrepreneurSignerEmail *string    `gorm:"column:entrepreneur_signer_email" json:"entrepreneur_signer_email"`
	StartupSignatureID      *string    `gorm:"column:startup_signature_id" json:"startup_signature_id"`
	StartupSignerEmail      *string    `gorm:"column:startup_signer_email" json:"startup_signer_email"`
	AgreementFile           *string    `gorm:"column:agreement_file" json:"agreement_file"`
	EquityAwardFile         *string    `gorm:"column:equity_award_file" json:"equity_award_file"`
}

// TableName Applicant's table name
func (*ApplicantDto) TableName() string {
	return TableNameApplicant
}
