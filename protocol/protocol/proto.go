package protocol

import (
	"backend/shared"
	"encoding/json"
	"errors"
	"slices"
)

type Protocol struct {
	id    int     `json:"id" db:"id" example:"123"`
	Name  string  `json:"name" db:"name" example:"SAAS Company"`
	Tasks []*Task `json:"tasks"`
}

func (p *Protocol) Id() int {
	return p.id
}

func (p *Protocol) Roles() []shared.StartupRole {
	var result []shared.StartupRole
	for idx := range p.Tasks {
		if !slices.Contains(result, p.Tasks[idx].StartupRole) && p.Tasks[idx].StartupRole != "" {
			result = append(result, p.Tasks[idx].StartupRole)
		}
	}
	return result
}

func (p *Protocol) Total() int {
	return len(p.Tasks)
}

func (p *Protocol) SampleTasks() []*Task {
	var resp []*Task
	for i, t := range p.Tasks {
		if i > 3 {
			break
		}
		if len(t.DependsOn) == 0 {
			resp = append(resp, t)
		}
	}
	return resp
}

func ToDto(p *Protocol) *ProtocolDto {
	var tasks []*TaskDto
	var r string
	for _, t := range p.Tasks {
		r = string(t.StartupRole)

		tasks = append(tasks, &TaskDto{
			ID:         t.Id,
			ProtocolID: p.id,
			Role:       r,
			CompleteFn: t.CompleteFn,
			DependsOn:  t.DependsOn, // todo use gorm serializer:json
			Title:      t.Title,
			Content:    t.Content,
		})
	}
	return &ProtocolDto{
		ID:    p.id,
		Name:  p.Name,
		Tasks: tasks,
	}
}

func parseDeps(src string) (DepList, error) {
	var result DepList
	err := json.Unmarshal([]byte(src), &result)
	if err != nil {
		return result, errors.New("unable to decode DepList: '" + src + "'")
	}
	return result, nil
}

func encodeDeps(d DepList) *string {
	var s string
	b, _ := json.Marshal(d)
	s = string(b)
	return &s
}

func FromDto(dto *ProtocolDto) *Protocol {
	var tasks []*Task
	for _, t := range dto.Tasks {
		tasks = append(tasks, &Task{
			Id:          t.ID,
			StartupRole: shared.StartupRole(t.Role),
			Title:       t.Title,
			Content:     t.Content,
			DependsOn:   t.DependsOn,
			CompleteFn:  "",
		})
	}
	return &Protocol{
		id:    dto.ID,
		Name:  dto.Name,
		Tasks: tasks,
	}
}
