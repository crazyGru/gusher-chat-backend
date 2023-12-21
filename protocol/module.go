package protocol

import (
	"backend/protocol/progress"
	"backend/protocol/protocol"
	"errors"
)

var ErrTaskNotInStartupProtocol = errors.New("task does not belong to startup protocol")
var ErrDependenciesNotComplete = errors.New("one or more dependent tasks are not complete")
var ErrStartupProtocolNotFound = errors.New("could not find startup protocol with specified id")

type ProtocolRepository interface {
	Find(params ProtocolSearch) ([]*protocol.Protocol, error)
	Count(params ProtocolSearch) (int, error)
	Save(protocol *protocol.Protocol) error
	Delete(protocol *protocol.Protocol)
	FindStartupProtocol(startupId int, protocolId int) (*StartupProtocol, error)
	SaveStartupProtocol(startupProtocol *StartupProtocol) error
}

type ProtocolTaskStartupRepository interface {
	Find(startupId int) ([]int, error)
	Save(startupId int, tasks []int)
}

type Module struct {
	repo ProtocolRepository
	ptsr ProtocolTaskStartupRepository
}

func NewModule(repo ProtocolRepository, ptsr ProtocolTaskStartupRepository) *Module {
	return &Module{repo: repo, ptsr: ptsr}
}

func (m *Module) Find(params ProtocolSearch) ([]*protocol.Protocol, error) {
	return m.repo.Find(params)
}

func (m *Module) Save(protocol *protocol.Protocol) error {
	return m.repo.Save(protocol)
}

func (m *Module) FindStartupProtocol(startupId int, protocolId int) (*StartupProtocol, error) {
	return m.repo.FindStartupProtocol(startupId, protocolId)
}

func (m *Module) GetProgress(startupId int, protocolId int) (*progress.Progress, error) {
	idList, err := m.ptsr.Find(startupId)
	if err != nil {
		return nil, err
	}
	protocols, err := m.Find(ProtocolSearch{Id: protocolId})
	if err != nil {
		return nil, err
	}
	return progress.NewProgress(startupId, protocols[0], idList), nil
}

func (m *Module) SaveProgress(p *progress.Progress) error {
	m.ptsr.Save(p.StartupId(), p.CompleteTaskIds())
	return nil
}

func (m *Module) SaveStartupProtocol(startupProtocol *StartupProtocol) error {
	return m.repo.SaveStartupProtocol(startupProtocol)
}
