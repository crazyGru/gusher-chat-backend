package role

import (
	"backend/shared"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"time"
)

type Role struct {
	id                int `db:"id"`
	protocolRole      []shared.StartupRole
	startup           int
	title             string
	description       string
	equity            float64
	applicants        int
	selectedApplicant int
	views             int
	status            string
	vanity            string
	created           time.Time
	modified          time.Time
	opened            time.Time
	closed            time.Time
	filled            time.Time
	awarded           time.Time
	blocked           time.Time
	hidden            time.Time
	featured          bool
	goals             []*Goal
}

func (r *Role) Id() int {
	return r.id
}

func (r *Role) ProtocolRole() []shared.StartupRole {
	return r.protocolRole
}

func (r *Role) SetProtocolRole(pr []shared.StartupRole) {
	r.protocolRole = pr
}

func FromDto(dto *RoleDto) *Role {
	var protocolRoles []shared.StartupRole
	if dto.ProtocolRole != "" {
		err := json.Unmarshal([]byte(dto.ProtocolRole), &protocolRoles)
		if err != nil {
			log.Errorf("role id=%v, failed to parse JSON for protocol_role", dto.ID)
		}
	}
	var goals []*Goal
	for _, g := range dto.Goals {
		goals = append(goals, goalFromDto(g))
	}
	return &Role{
		id:                dto.ID,
		protocolRole:      protocolRoles,
		startup:           dto.Startup,
		title:             dto.Title,
		description:       dto.Description,
		equity:            dto.Equity,
		applicants:        dto.Applicants,
		selectedApplicant: dto.SelectedApplicant,
		views:             dto.Views,
		status:            dto.Status,
		vanity:            dto.Vanity,
		created:           dto.Created,
		modified:          dto.Modified,
		opened:            dto.Opened,
		closed:            dto.Closed,
		filled:            dto.Filled,
		awarded:           dto.Awarded,
		blocked:           dto.Blocked,
		hidden:            dto.Hidden,
		featured:          dto.Featured == 1,
		goals:             goals,
	}
}

func ToDto(r *Role) (roleDto *RoleDto) {
	var featured int
	if r.featured {
		featured = 1
	} else {
		featured = 0
	}
	pr, _ := json.Marshal(r.protocolRole)
	var goals []*GoalDto
	for _, g := range r.goals {
		goals = append(goals, goalToDto(g))
	}
	return &RoleDto{
		ID:                r.id,
		Startup:           r.startup,
		Title:             r.title,
		Description:       r.description,
		Equity:            r.equity,
		Applicants:        r.applicants,
		SelectedApplicant: r.selectedApplicant,
		Views:             r.views,
		Status:            r.status,
		Vanity:            r.vanity,
		Created:           r.created,
		Modified:          r.modified,
		Opened:            r.opened,
		Closed:            r.closed,
		Filled:            r.filled,
		Awarded:           r.awarded,
		Blocked:           r.blocked,
		Hidden:            r.hidden,
		Featured:          featured,
		ProtocolRole:      string(pr),
		Goals:             goals,
	}
}

func goalFromDto(g *GoalDto) *Goal {
	return &Goal{
		id:        g.ID,
		parent:    g.Parent,
		startup:   g.Startup,
		role:      g.Role,
		user:      g.User,
		title:     g.Title,
		status:    g.Status,
		created:   g.Created,
		modified:  g.Modified,
		activated: g.Activated,
		reached:   g.Reached,
		approved:  g.Approved,
		blocked:   g.Blocked,
		hidden:    g.Hidden,
		notified:  g.Notified == 1,
	}
}

func goalToDto(g *Goal) *GoalDto {
	var notified int
	if g.notified {
		notified = 1
	} else {
		notified = 0
	}
	return &GoalDto{
		ID:        g.id,
		Parent:    g.parent,
		Startup:   g.startup,
		Role:      g.role,
		User:      g.user,
		Title:     g.title,
		Status:    g.status,
		Created:   g.created,
		Modified:  g.modified,
		Activated: g.activated,
		Reached:   g.reached,
		Approved:  g.approved,
		Blocked:   g.blocked,
		Hidden:    g.hidden,
		Notified:  notified,
	}
}
