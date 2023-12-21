package block

import "time"

type UserBlock struct {
	UserId    int
	BlockedId int
	Date      time.Time
	Note      string
}

func FromDto(dto *UserBlockDto) *UserBlock {
	return &UserBlock{
		UserId:    dto.UserID,
		BlockedId: dto.BlockedID,
		Date:      dto.Date,
		Note:      dto.Note,
	}
}

func ToDto(u *UserBlock) *UserBlockDto {
	return &UserBlockDto{
		UserID:    u.UserId,
		BlockedID: u.BlockedId,
		Date:      u.Date,
		Note:      u.Note,
	}
}
