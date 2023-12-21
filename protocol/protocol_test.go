package protocol

import (
	protocol2 "backend/protocol/protocol"
	"errors"
	"slices"
	"testing"
)

func setupStruct() *StartupProtocol {
	protocol := &protocol2.Protocol{
		Tasks: []*protocol2.Task{
			&protocol2.Task{
				Id:          1,
				StartupRole: "owner",
				Title:       "Do Something",
				Content:     "Do something already",
				DependsOn:   nil,
				CompleteFn:  "",
			},
			&protocol2.Task{
				Id:          2,
				StartupRole: "owner",
				Title:       "Do Something",
				Content:     "Do something already",
				DependsOn:   nil,
				CompleteFn:  "",
			},
			&protocol2.Task{
				Id:          3,
				StartupRole: "owner",
				Title:       "Do Something",
				Content:     "Do something already",
				DependsOn:   []int{1, 2},
				CompleteFn:  "",
			},
			&protocol2.Task{
				Id:          4,
				StartupRole: "owner",
				Title:       "Do Something",
				Content:     "Do something already",
				DependsOn:   nil,
				CompleteFn:  "",
			},
		},
	}
	startupProtocol := &StartupProtocol{
		StartupId:     0,
		Protocol:      protocol,
		CompleteTasks: []int{1},
	}
	return startupProtocol
}

func TestStartupProtocol_Complete(t *testing.T) {
	var err error
	sp := setupStruct()
	err = sp.Complete(3)
	if !errors.Is(err, ErrDependenciesNotComplete) {
		t.Errorf("Should not be able to complete task if dependencies not complete")
	}
	sp.Complete(2)
	err = sp.Complete(3)
	if err != nil {
		t.Errorf("Should be able to complete task if dependencies complete")
	}
	sp.Complete(4)
	if err != nil {
		t.Errorf("Should be able to complete task if no dependencies")
	}
	sp.Complete(256)
	if slices.Contains(sp.CompleteTasks, 256) {
		t.Errorf("Tasks that are not in protocol can not be completed")
	}
}

func TestStartupProtocol_NotComplete(t *testing.T) {
	sp := setupStruct()
	sp.NotComplete(1)
	if slices.Contains(sp.CompleteTasks, 1) {
		t.Errorf("Failed to mark task as not complete")
	}
}

func TestStartupProtocol_canComplete(t *testing.T) {

}

func TestStartupProtocol_taskById(t *testing.T) {
}
