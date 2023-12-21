package progress

import (
	"backend/protocol/protocol"
	"backend/shared"
	"errors"
	"slices"
)

var ErrTaskNotInStartupProtocol = errors.New("task does not belong to startup protocol")
var ErrDependenciesNotComplete = errors.New("one or more dependent tasks are not complete")
var ErrStartupProtocolNotFound = errors.New("could not find startup protocol with specified id")

type Progress struct {
	startupId     int
	completeTasks []int
	protocol      *protocol.Protocol
}

func NewProgress(startupId int, p *protocol.Protocol, completeTasks []int) *Progress {
	return &Progress{
		startupId:     startupId,
		completeTasks: completeTasks,
		protocol:      p,
	}
}

func (p *Progress) CompleteTaskIds() []int {
	return p.completeTasks
}

func (p *Progress) StartupId() int {
	return p.startupId
}

func (p *Progress) ProtocolName() string {
	return p.protocol.Name
}

func (p *Progress) ProtocolId() int {
	return p.protocol.Id()
}

func (p *Progress) Complete(taskId int) error {
	if p.taskById(taskId) == nil {
		return ErrTaskNotInStartupProtocol
	}
	if !p.canComplete(taskId) {
		return ErrDependenciesNotComplete
	}
	if !slices.Contains(p.completeTasks, taskId) {
		p.completeTasks = append(p.completeTasks, taskId)
	}
	return nil
}

func (p *Progress) NotComplete(taskId int) error {
	if p.taskById(taskId) == nil {
		return ErrTaskNotInStartupProtocol
	}
	p.completeTasks = slices.DeleteFunc(p.completeTasks, func(i int) bool {
		return i == taskId
	})
	return nil
}

func (p *Progress) canComplete(taskId int) bool {
	task := p.taskById(taskId)
	if task == nil {
		return false
	}
	if slices.Contains(p.completeTasks, taskId) {
		return false
	}
	for key := range task.DependsOn {
		if !slices.Contains(p.completeTasks, task.DependsOn[key]) {
			return false
		}
	}
	return true
}

func (p *Progress) taskById(id int) *protocol.Task {
	i := slices.IndexFunc(p.protocol.Tasks, func(task *protocol.Task) bool {
		return task.Id == id
	})
	if i != -1 {
		return p.protocol.Tasks[i]
	}
	return nil
}

func (p *Progress) RoleTasks(startupRoles []shared.StartupRole) []*protocol.Task {
	var result []*protocol.Task
	for k := range p.protocol.Tasks {
		if slices.Contains(startupRoles, p.protocol.Tasks[k].StartupRole) {
			result = append(result, p.protocol.Tasks[k])
		}
	}
	return result

}

func (p *Progress) PendingTasks(startupRoles []shared.StartupRole) []*protocol.Task {
	var result []*protocol.Task
	for k := range p.protocol.Tasks {
		if slices.Contains(startupRoles, "founder") || (slices.Contains(startupRoles, p.protocol.Tasks[k].StartupRole) && p.canComplete(p.protocol.Tasks[k].Id)) {
			result = append(result, p.protocol.Tasks[k])
		}
	}
	return result
}

func (p *Progress) CompleteTasks() []*protocol.Task {
	var result []*protocol.Task
	for k := range p.protocol.Tasks {
		if slices.Contains(p.completeTasks, p.protocol.Tasks[k].Id) {
			result = append(result, p.protocol.Tasks[k])
		}
	}
	return result
}

func (p *Progress) TotalTasks() int {
	return len(p.protocol.Tasks)
}
