package repo

import (
	"backend/chat/room"

	"gorm.io/gorm"
)

type RoomRepository struct {
	db *gorm.DB
}

// SaveRoom implements chat.RoomRepository.
func (*RoomRepository) SaveRoom(r *room.Room) {
	panic("unimplemented")
}

func NewRoomRepository(db *gorm.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

func (r *RoomRepository) CreateRoom(room room.Room) (*room.Room, error) {
	err := r.db.Create(&room).Error
	return &room, err
}
