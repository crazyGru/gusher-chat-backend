package protocol

import (
	"backend/protocol/protocol"
	"backend/shared"
	"slices"
)

type StartupProtocol struct {
	StartupId     int
	Protocol      *protocol.Protocol
	CompleteTasks []int `example:"1,2,3,4"` // IDs of complete tasks
}

func (p *StartupProtocol) Complete(taskId int) error {
	if p.taskById(taskId) == nil {
		return ErrTaskNotInStartupProtocol
	}
	if !p.canComplete(taskId) {
		return ErrDependenciesNotComplete
	}
	if !slices.Contains(p.CompleteTasks, taskId) {
		p.CompleteTasks = append(p.CompleteTasks, taskId)
	}
	return nil
}

func (p *StartupProtocol) NotComplete(taskId int) error {
	if p.taskById(taskId) == nil {
		return ErrTaskNotInStartupProtocol
	}
	p.CompleteTasks = slices.DeleteFunc(p.CompleteTasks, func(i int) bool {
		return i == taskId
	})
	return nil
}

func (p *StartupProtocol) canComplete(taskId int) bool {
	task := p.taskById(taskId)
	if task == nil {
		return false
	}
	for key := range task.DependsOn {
		if !slices.Contains(p.CompleteTasks, task.DependsOn[key]) {
			return false
		}
	}
	return true
}

func (p *StartupProtocol) taskById(id int) *protocol.Task {
	i := slices.IndexFunc(p.Protocol.Tasks, func(task *protocol.Task) bool {
		return task.Id == id
	})
	if i != -1 {
		return p.Protocol.Tasks[i]
	}
	return nil
}

func (p *StartupProtocol) RoleTasks(startupRoles []shared.StartupRole) []*protocol.Task {
	var result []*protocol.Task
	for k := range p.Protocol.Tasks {
		if slices.Contains(startupRoles, p.Protocol.Tasks[k].StartupRole) {
			result = append(result, p.Protocol.Tasks[k])
		}
	}
	return result

}

func (p *StartupProtocol) PendingTasks(startupRoles []shared.StartupRole) []*protocol.Task {
	var result []*protocol.Task
	for k := range p.Protocol.Tasks {
		if slices.Contains(startupRoles, p.Protocol.Tasks[k].StartupRole) && p.canComplete(p.Protocol.Tasks[k].Id) {
			result = append(result, p.Protocol.Tasks[k])
		}
	}
	return result
}
