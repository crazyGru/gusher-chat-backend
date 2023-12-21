package startup

import (
	"time"
)

type Startup struct {
	id                int
	user              int
	protocol          int
	status            string
	userTitle         string
	userDescription   string
	userEquity        float64
	name              string
	title             string
	description       string
	avatar            Image
	logo              Image
	icon              Image
	cover             Image
	video             Video
	pitch             string
	equityMin         float64
	equityTotal       float64
	deadline          int
	vanity            string
	companyType       string
	companyState      string
	followers         int
	applicants        int
	views             int
	reports           int
	created           time.Time
	modified          time.Time
	submitted         time.Time
	reviewed          time.Time
	signed            time.Time
	published         time.Time
	development       time.Time
	finished          time.Time
	launched          time.Time
	awardUploaded     time.Time
	unpublished       time.Time
	hidden            time.Time
	blocked           time.Time
	privacy           string
	confidential      bool
	featured          bool
	notified          bool
	address           Address
	signature         Signature
	incorporationFile string
	equityAwardFile   string
	topics            []*Topic
}

type Image struct {
	Src     string `json:"src"`
	Resized string `json:"resized"`
}

type Video struct {
	Src    string `json:"src"`
	Src720 string `json:"src720"`
	Job    string `json:"job"`
}

type Address struct {
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	AddressLine3 string `json:"address_line_3"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zipcode      string `json:"zipcode"`
	Country      string `json:"country"`
	Email        string `json:"email"`
	Attn         string `json:"attn"`
}

type Signature struct {
	SignatureRequest string `json:"signature_request"`
	SignatureID      string `json:"signature_id"`
	AgreementFile    string `json:"agreement_file"`
}

func (s *Startup) Id() int                   { return s.id }
func (s *Startup) User() int                 { return s.user }
func (s *Startup) Protocol() int             { return s.protocol }
func (s *Startup) Status() string            { return s.status }
func (s *Startup) UserTitle() string         { return s.userTitle }
func (s *Startup) UserDescription() string   { return s.userDescription }
func (s *Startup) UserEquity() float64       { return s.userEquity }
func (s *Startup) Name() string              { return s.name }
func (s *Startup) Title() string             { return s.title }
func (s *Startup) Description() string       { return s.description }
func (s *Startup) Avatar() Image             { return s.avatar }
func (s *Startup) Logo() Image               { return s.logo }
func (s *Startup) Icon() Image               { return s.icon }
func (s *Startup) Cover() Image              { return s.cover }
func (s *Startup) Video() Video              { return s.video }
func (s *Startup) Pitch() string             { return s.pitch }
func (s *Startup) EquityMin() float64        { return s.equityMin }
func (s *Startup) EquityTotal() float64      { return s.equityTotal }
func (s *Startup) Deadline() int             { return s.deadline }
func (s *Startup) Vanity() string            { return s.vanity }
func (s *Startup) CompanyType() string       { return s.companyType }
func (s *Startup) CompanyState() string      { return s.companyState }
func (s *Startup) Followers() int            { return s.followers }
func (s *Startup) Applicants() int           { return s.applicants }
func (s *Startup) Views() int                { return s.views }
func (s *Startup) Reports() int              { return s.reports }
func (s *Startup) Created() time.Time        { return s.created }
func (s *Startup) Modified() time.Time       { return s.modified }
func (s *Startup) Submitted() time.Time      { return s.submitted }
func (s *Startup) Reviewed() time.Time       { return s.reviewed }
func (s *Startup) Signed() time.Time         { return s.signed }
func (s *Startup) Published() time.Time      { return s.published }
func (s *Startup) Development() time.Time    { return s.development }
func (s *Startup) Finished() time.Time       { return s.finished }
func (s *Startup) Launched() time.Time       { return s.launched }
func (s *Startup) AwardUploaded() time.Time  { return s.awardUploaded }
func (s *Startup) Unpublished() time.Time    { return s.unpublished }
func (s *Startup) Hidden() time.Time         { return s.hidden }
func (s *Startup) Blocked() time.Time        { return s.blocked }
func (s *Startup) Privacy() string           { return s.privacy }
func (s *Startup) Confidential() bool        { return s.confidential }
func (s *Startup) Featured() bool            { return s.featured }
func (s *Startup) Notified() bool            { return s.notified }
func (s *Startup) Address() Address          { return s.address }
func (s *Startup) Signature() Signature      { return s.signature }
func (s *Startup) IncorporationFile() string { return s.incorporationFile }
func (s *Startup) EquityAwardFile() string   { return s.equityAwardFile }
func (s *Startup) Topics() []*Topic          { return s.topics }

func (s *Startup) SetProtocol(protocolId int) {
	s.protocol = protocolId
}

func boolToInt(val bool) int {
	if val {
		return 1
	}
	return 0
}

func FromDto(dto *StartupDto) *Startup {
	return &Startup{
		id:              dto.ID,
		user:            dto.User,
		protocol:        dto.Protocol,
		status:          dto.Status,
		userTitle:       dto.UserTitle,
		userDescription: dto.UserDescription,
		userEquity:      dto.UserEquity,
		name:            dto.Name,
		title:           dto.Title,
		description:     dto.Description,
		avatar: Image{
			Src:     dto.Avatar,
			Resized: dto.AvatarMd,
		},
		logo: Image{
			Src:     dto.Logo,
			Resized: "",
		},
		icon: Image{
			Src:     dto.Icon,
			Resized: "",
		},
		cover: Image{
			Src:     dto.Cover,
			Resized: dto.CoverMd,
		},
		video: Video{
			Src:    dto.Video,
			Src720: dto.Video720URL,
			Job:    dto.Video720Job,
		},
		pitch:         dto.Pitch,
		equityMin:     dto.EquityMin,
		equityTotal:   dto.EquityTotal,
		deadline:      dto.Deadline,
		vanity:        dto.Vanity,
		companyType:   dto.CompanyType,
		companyState:  dto.CompanyState,
		followers:     dto.Followers,
		applicants:    dto.Applicants,
		views:         dto.Views,
		reports:       dto.Reports,
		created:       dto.Created,
		modified:      dto.Modified,
		submitted:     dto.Submitted,
		reviewed:      dto.Reviewed,
		signed:        dto.Signed,
		published:     dto.Published,
		development:   dto.Development,
		finished:      dto.Finished,
		launched:      dto.Launched,
		awardUploaded: dto.AwardUploaded,
		unpublished:   dto.Unpublished,
		hidden:        dto.Hidden,
		blocked:       dto.Blocked,
		privacy:       dto.Privacy,
		confidential:  dto.Confidential == 1,
		featured:      dto.Featured == 1,
		notified:      dto.Notified == 1,
		address: Address{
			AddressLine1: dto.AddressLine1,
			AddressLine2: dto.AddressLine2,
			AddressLine3: dto.AddressLine3,
			City:         dto.City,
			State:        dto.State,
			Zipcode:      dto.Zipcode,
			Country:      dto.Country,
			Email:        dto.Email,
			Attn:         dto.Attn,
		},
		signature: Signature{
			SignatureRequest: dto.SignatureRequest,
			SignatureID:      dto.SignatureID,
			AgreementFile:    dto.AgreementFile,
		},
		incorporationFile: dto.IncorporationFile,
		equityAwardFile:   dto.EquityAwardFile,
		topics:            dto.Topics,
	}
}

func ToDto(s *Startup) *StartupDto {
	return &StartupDto{
		ID:                s.id,
		User:              s.user,
		UserTitle:         s.userTitle,
		UserDescription:   s.userDescription,
		UserEquity:        s.userEquity,
		Name:              s.name,
		Title:             s.title,
		Description:       s.description,
		Avatar:            s.avatar.Src,
		AvatarMd:          s.avatar.Resized,
		Logo:              s.logo.Src,
		Icon:              s.icon.Src,
		Cover:             s.cover.Src,
		CoverMd:           s.cover.Resized,
		Video:             s.video.Src,
		Video720URL:       s.video.Src720,
		Video720Job:       s.video.Job,
		Pitch:             s.pitch,
		EquityMin:         s.equityMin,
		EquityTotal:       s.equityTotal,
		Deadline:          s.deadline,
		Vanity:            s.vanity,
		CompanyType:       s.companyType,
		CompanyState:      s.companyState,
		Followers:         s.followers,
		Applicants:        s.applicants,
		Views:             s.views,
		Reports:           s.reports,
		Status:            s.status,
		Created:           s.created,
		Modified:          s.modified,
		Submitted:         s.submitted,
		Reviewed:          s.reviewed,
		Signed:            s.signed,
		Published:         s.published,
		Development:       s.development,
		Finished:          s.finished,
		Launched:          s.launched,
		AwardUploaded:     s.awardUploaded,
		Unpublished:       s.unpublished,
		Hidden:            s.hidden,
		Blocked:           s.blocked,
		Privacy:           s.privacy,
		Confidential:      boolToInt(s.confidential),
		Featured:          boolToInt(s.featured),
		Notified:          boolToInt(s.notified),
		AddressLine1:      s.address.AddressLine1,
		AddressLine2:      s.address.AddressLine2,
		AddressLine3:      s.address.AddressLine3,
		City:              s.address.City,
		State:             s.address.State,
		Zipcode:           s.address.Zipcode,
		Country:           s.address.Country,
		Email:             s.address.Email,
		Attn:              s.address.Attn,
		SignatureRequest:  s.signature.SignatureRequest,
		SignatureID:       s.signature.SignatureID,
		AgreementFile:     s.signature.AgreementFile,
		IncorporationFile: s.incorporationFile,
		EquityAwardFile:   s.equityAwardFile,
		Protocol:          s.protocol,
		Topics:            s.topics,
	}
}
