package chat

import (
	"backend/chat/message"
	"backend/chat/room"
	"fmt"
	"time"
)

type Module struct {
	roomRepository    RoomRepository
	memberRepository  MemberRepository
	messageRepository MessageRepository
}

func NewModule(roomRepo RoomRepository, memberRepo MemberRepository, messageRepo MessageRepository) *Module {
	return &Module{
		roomRepository:    roomRepo,
		memberRepository:  memberRepo,
		messageRepository: messageRepo,
	}
}

// func (m *Module) Find(params MessageSearch) ([]*chat.Message, error) {
// 	return nil
// 	// return m.roomRepository.Find(params)
// }

func (m *Module) CreateRoom(superAdmin int, name string) (*room.Room, error) {
	room := room.Room{
		Admins: superAdmin,
		Name:   name,
	}
	return m.roomRepository.CreateRoom(room)
}

func (m *Module) GetMessageByRoomID(roomID uint) ([]*message.Message, error) {
	return m.messageRepository.FindByRoomID(roomID)
}

func (m *Module) GetRoomsByUser(userID int) ([]int, error) {
	return m.memberRepository.GetRoomsByUserID(userID)
}

func (m *Module) InviteMember(userID, roomID uint) error {
	member := room.RoomMember{
		RoomID:    int(roomID),
		UserID:    int(userID),
		Activated: time.Now(),
	}

	return m.memberRepository.CreateMember(member)
}

func (m *Module) RemoveMember(userID, roomID uint) error {
	member := room.RoomMember{
		RoomID: int(roomID),
		UserID: int(userID),
	}

	return m.memberRepository.RemoveMember(member)
}

func (m *Module) SendMessage(userID, roomID uint, content string) error {
	message := message.Message{
		RoomID:    int(roomID),
		SenderID:  int(userID),
		CreatedAt: time.Now(),
		TextBody:  content,
	}

	if err := m.messageRepository.CreateMessage(&message); err != nil {
		return fmt.Errorf("failed to save message: %w", err)
	}
	return nil
}
