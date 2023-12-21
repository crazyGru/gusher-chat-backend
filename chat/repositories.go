package chat

import (
	"backend/chat/message"
	"backend/chat/room"
)

type ChatRepository interface {
}

type MessageRepository interface {
	CreateMessage(message *message.Message) error
	FindByRoomID(roomID uint) ([]*message.Message, error)
}

type RoomRepository interface {
	CreateRoom(r room.Room) (*room.Room, error)
	SaveRoom(r *room.Room)
}

type MemberRepository interface {
	CreateMember(m room.RoomMember) error
	RemoveMember(m room.RoomMember) error
	GetRoomsByUserID(userID int) ([]int, error)
}
