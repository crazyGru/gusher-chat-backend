package repo

import (
	"backend/chat/chat"

	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (m *ChatRepository) Delete(item *chat.Chat) error {
	//TODO
	return nil
}
