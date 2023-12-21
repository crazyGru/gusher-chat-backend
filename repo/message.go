package repo

import (
	"backend/chat/message"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) CreateMessage(m message.Message) error {
	err := r.db.Create(&m).Error
	return err
}
func (r MessageRepository) FindByRoomID(roomID uint) ([]message.Message, error) {
	var messages []message.Message
	err := r.db.Where("channel = ?", roomID).Find(&messages).Error
	return messages, err
}
