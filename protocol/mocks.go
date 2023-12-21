package protocol

import (
	"backend/protocol/protocol"
	"slices"
)

type MockRepository struct {
	Protocols        []*protocol.Protocol
	StartupProtocols []*StartupProtocol
}

func (m *MockRepository) Count(params ProtocolSearch) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) Delete(protocol *protocol.Protocol) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) mockProtocol(name string) *protocol.Protocol {
	return protocol.FromDto(&protocol.ProtocolDto{
		ID:    123,
		Tasks: nil,
	})
}

func (m *MockRepository) mockTasks() []*protocol.Task {
	var result []*protocol.Task
	result = append(result, &protocol.Task{
		Id:          1,
		StartupRole: "founder",
		Title:       "Do Something",
		Content:     "Do something already",
		DependsOn:   []int{},
		CompleteFn:  "",
	})
	result = append(result, &protocol.Task{
		Id:          2,
		StartupRole: "founder",
		Title:       "Do Something",
		Content:     "Do something already",
		DependsOn:   []int{},
		CompleteFn:  "",
	})
	result = append(result, &protocol.Task{
		Id:          3,
		StartupRole: "founder",
		Title:       "Do Something",
		Content:     "Do something already",
		DependsOn:   []int{},
		CompleteFn:  "",
	})
	result = append(result, &protocol.Task{
		Id:          4,
		StartupRole: "founder",
		Title:       "Do Something",
		Content:     "Do something already",
		DependsOn:   []int{1, 2, 3},
		CompleteFn:  "",
	})
	return result

}

func (m *MockRepository) Find(params ProtocolSearch) ([]*protocol.Protocol, error) {
	var result []*protocol.Protocol
	if params.Id != 0 {
		result = slices.DeleteFunc(m.Protocols, func(protocol *protocol.Protocol) bool {
			return protocol.Id() != params.Id
		})
	} else if params.Name != "" {
		result = slices.DeleteFunc(m.Protocols, func(protocol *protocol.Protocol) bool {
			return protocol.Name != params.Name
		})
	} else {
		result = m.Protocols
	}

	return result, nil
}

func (m *MockRepository) Save(p *protocol.Protocol) error {
	index := slices.IndexFunc(m.Protocols, func(pr *protocol.Protocol) bool {
		return p.Id() == pr.Id()
	})
	if index != -1 {
		m.Protocols[index] = p
	} else {
		m.Protocols = append(m.Protocols, p)
		//if p.Id() == 0 {
		//	p.Id() = len(m.Protocols)
		//}
	}
	return nil
}

func (m *MockRepository) FindStartupProtocol(startupId int, protocolId int) (*StartupProtocol, error) {
	index := slices.IndexFunc(m.StartupProtocols, func(protocol *StartupProtocol) bool {
		return protocol.StartupId == startupId
	})
	if index == -1 {
		return nil, ErrStartupProtocolNotFound
	}
	return m.StartupProtocols[index], nil
}

func (m *MockRepository) SaveStartupProtocol(startupProtocol *StartupProtocol) error {
	index := slices.IndexFunc(m.StartupProtocols, func(sp *StartupProtocol) bool {
		return sp.StartupId == startupProtocol.StartupId
	})
	if index != -1 {
		m.StartupProtocols[index] = startupProtocol
	} else {
		m.StartupProtocols = append(m.StartupProtocols, startupProtocol)
	}
	return nil
}

type MockPtsrRepo struct {
}

func (m MockPtsrRepo) Find(startupId int) ([]int, error) {
	return []int{1, 2, 3, 4}, nil
}

func (m MockPtsrRepo) Save(startupId int, tasks []int) {
}
