package room

import "time"

type Room struct {
	Id           int       `db:"id"`
	Startup      int       `db:"startup"`
	Name         string    `db:"name"`
	Privacy      string    `db:"privacy"`
	Direct       int       `db:"direct"`
	Admins       int       `db:"admins"`
	Recruiters   int       `db:"recruiters"`
	Broadcasters int       `db:"broadcasters"`
	Status       string    `db:"status"`
	Created      time.Time `db:"created"`
	Archived     time.Time `db:"archived"`
	Hidden       time.Time `db:"hidden"`
	Blocked      time.Time `db:"blocked"`
}

type RoomMember struct {
	RoomID     int       `db:"channel"`
	UserID     int       `db:"user"`
	Status     bool      `db:"status"`
	Activated  time.Time `db:"activated"`
	Blocked    time.Time `db:"blocked"`
	IsAdmin    bool      `db:"admin"`
	IsNotified bool      `db:"notified"`
}

func (r *Room) GetId() int               { return r.Id }
func (r *Room) GetStartup() int          { return r.Startup }
func (r *Room) GetName() string          { return r.Name }
func (r *Room) GetPrivacy() string       { return r.Privacy }
func (r *Room) GetDirect() int           { return r.Direct }
func (r *Room) GetAdmins() int           { return r.Admins }
func (r *Room) GetRecruiters() int       { return r.Recruiters }
func (r *Room) GetBroadcasters() int     { return r.Broadcasters }
func (r *Room) GetStatus() string        { return r.Status }
func (r *Room) GetCreatedAt() time.Time  { return r.Created }
func (r *Room) GetArchivedAt() time.Time { return r.Archived }
func (r *Room) GetHiddenAt() time.Time   { return r.Hidden }
func (r *Room) GetBlockedAt() time.Time  { return r.Blocked }

func FromDto(dto *RoomDto) *Room {
	return &Room{
		Id:           dto.Id,
		Startup:      dto.Startup,
		Name:         dto.Name,
		Privacy:      dto.Privacy,
		Direct:       dto.Direct,
		Admins:       dto.Admins,
		Recruiters:   dto.Recruiters,
		Broadcasters: dto.Broadcasters,
		Status:       dto.Status,
		Created:      dto.Created,
		Archived:     dto.Archived,
		Hidden:       dto.Hidden,
		Blocked:      dto.Blocked,
	}
}

func ToDto(r *Room) *RoomDto {
	return &RoomDto{
		Id:           r.Id,
		Startup:      r.Startup,
		Name:         r.Name,
		Privacy:      r.Privacy,
		Direct:       r.Direct,
		Admins:       r.Admins,
		Recruiters:   r.Recruiters,
		Broadcasters: r.Broadcasters,
		Status:       r.Status,
		Created:      r.Created,
		Archived:     r.Archived,
		Hidden:       r.Hidden,
		Blocked:      r.Blocked,
	}
}
