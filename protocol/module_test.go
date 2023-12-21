package protocol

import (
	"backend/protocol/protocol"
	"reflect"
	"testing"
)

func mockPtsrRepo() ProtocolTaskStartupRepository {
	return &MockPtsrRepo{}
}

func mockRepo() ProtocolRepository {
	//tasks := []*protocol.Task{
	//	&protocol.Task{
	//		Id:          1,
	//		StartupRole: "owner",
	//		Title:       "Do Something",
	//		Content:     "Do something already",
	//		DependsOn:   nil,
	//	},
	//	&protocol.Task{
	//		Id:          2,
	//		StartupRole: "owner",
	//		Title:       "Do Something",
	//		Content:     "Do something already",
	//		DependsOn:   nil,
	//	},
	//	&protocol.Task{
	//		Id:          3,
	//		StartupRole: "owner",
	//		Title:       "Do Something",
	//		Content:     "Do something already",
	//		DependsOn:   []int{1, 3},
	//	},
	//}
	protocols := []*protocol.Protocol{
		protocol.FromDto(&protocol.ProtocolDto{
			ID:    1,
			Tasks: nil,
		}),
		protocol.FromDto(&protocol.ProtocolDto{
			ID:    2,
			Tasks: nil,
		}),
		protocol.FromDto(&protocol.ProtocolDto{
			ID:    3,
			Tasks: nil,
		}),
	}
	startupProtocols := []*StartupProtocol{
		&StartupProtocol{
			StartupId:     1,
			Protocol:      protocols[1],
			CompleteTasks: []int{1},
		},
		&StartupProtocol{
			StartupId:     2,
			Protocol:      protocols[1],
			CompleteTasks: []int{1, 2},
		},
	}

	return &MockRepository{
		Protocols:        protocols,
		StartupProtocols: startupProtocols,
	}
}

func TestNewModule(t *testing.T) {
	m := NewModule(mockRepo(), mockPtsrRepo())
	if reflect.TypeOf(m).String() != "*protocol.Module" {
		t.Errorf("Could not create module, expected *protocol.Module, got: %s", reflect.TypeOf(m).String())
	}
}

func TestModule_FindProtocols(t *testing.T) {
	m := NewModule(mockRepo(), mockPtsrRepo())
	protocols, err := m.Find(ProtocolSearch{})
	if err != nil {
		t.Errorf("expected list of protocols, got error: %s", err)
	}
	if len(protocols) == 0 {
		t.Errorf("empty protocols list returned")
	}

}

func TestModule_FindStartupProtocol(t *testing.T) {
	m := NewModule(mockRepo(), mockPtsrRepo())
	sp, err := m.FindStartupProtocol(1, 1)
	if err != nil {
		t.Errorf("expected startup protocol, got error: %s", err)
	}
	if sp.StartupId != 1 {
		t.Errorf("Incorrect startupId, expected %d got %d", 1, sp.StartupId)
	}
}

func TestModule_SaveProtocol(t *testing.T) {
	m := NewModule(mockRepo(), mockPtsrRepo())
	proto, _ := m.Find(ProtocolSearch{
		Id: 1,
	})
	err := m.Save(proto[0])
	if err != nil {
		t.Error(err)
	}

	newProto := protocol.FromDto(&protocol.ProtocolDto{
		ID:    123,
		Tasks: nil,
	})
	m.Save(newProto)
	res, _ := m.Find(ProtocolSearch{Id: 123})
	if res[0].Id() != 123 {
		t.Errorf("Could not save protocol")
	}
}

func TestModule_SaveStartupProtocol(t *testing.T) {
	m := NewModule(mockRepo(), mockPtsrRepo())
	sp, err := m.FindStartupProtocol(1, 1)
	if err != nil {
		t.Error(err)
	}
	err = m.SaveStartupProtocol(sp)
	if err != nil {
		t.Error(err)
	}

}
