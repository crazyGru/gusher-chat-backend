package access

import (
	"backend/chat"
	"backend/protocol"
	"backend/startup"
	"backend/user"
	"errors"
	"log"
)

var ErrAccessDenied = errors.New("access denied")

type Identity struct {
	UserId int `json:"id"`
}

type Module struct {
	startupModule  *startup.Module
	protocolModule *protocol.Module
	userModule     *user.Module
	idProvider     IdProvider
	chatModule     *chat.Module
}

func NewModule(idProvider IdProvider, startupModule *startup.Module, protocolModule *protocol.Module, userModule *user.Module, chatModule *chat.Module) *Module {
	return &Module{
		idProvider:     idProvider,
		startupModule:  startupModule,
		protocolModule: protocolModule,
		userModule:     userModule,
		chatModule:     chatModule,
	}
}

func (m *Module) GetIdentity(key string) (*Identity, error) {
	return m.idProvider.Get(key)
}

func (m *Module) CanAccessProtocolTask(userId int, startupId int, taskId int) bool {
	roles := m.startupModule.UserRolesInStartup(startupId, userId)
	protocolId, err := m.startupModule.StartupProtocolId(startupId)
	if err != nil {
		return false
	}
	sp, err := m.protocolModule.GetProgress(startupId, protocolId)
	if err != nil {
		return false
	}
	for _, task := range sp.RoleTasks(roles) {
		if task.Id == taskId {
			return true
		}
	}
	return false
}

func (m *Module) CanManageProtocols(userId int) bool {
	return m.isAdmin(userId)
}

func (m *Module) CanManageStartup(userId int, startupId int) bool {
	if m.isAdmin(userId) {
		return true
	}
	s, err := m.startupModule.FindStartups(startup.Search{
		Id: startupId,
	})
	if err != nil {
		return false
	}
	return s[0].User() == userId
}

func (m *Module) isAdmin(userId int) bool {
	u, err := m.userModule.FindUser(user.Search{Id: userId})
	if err != nil {
		log.Printf("user.isAdmin: %s", err.Error())
		return false
	}
	return u.Admin
}
