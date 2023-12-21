package repo

import (
	"backend/chat/room"

	"gorm.io/gorm"
)

type RoomMemberRepository struct {
	db *gorm.DB
}

// CreateMember implements chat.MemberRepository.
func (*RoomMemberRepository) CreateMember(m room.RoomMember) error {
	panic("unimplemented")
}

func NewRoomMemberRepository(db *gorm.DB) *RoomMemberRepository {
	return &RoomMemberRepository{db: db}
}

func (r *RoomMemberRepository) CreateRoomMember(m room.RoomMember) error {
	err := r.db.Create(&m).Error
	return err
}

func (r *RoomMemberRepository) RemoveMember(m room.RoomMember) error {

	// Use the Delete method with the struct to match both primary keys
	err := r.db.Where(&m).Delete(&room.RoomMember{}).Error
	return err
}

func (r *RoomMemberRepository) GetRoomsByUserID(userID int) ([]int, error) {
	var roomMembers []room.RoomMember
	err := r.db.Model(&room.RoomMember{}).Where("user_id = ?", userID).Find(&roomMembers).Error
	if err != nil {
		return nil, err
	}
	var roomIDs []int
	for _, member := range roomMembers {
		roomIDs = append(roomIDs, member.RoomID)
	}

	return roomIDs, nil
}
