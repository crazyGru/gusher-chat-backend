package chat

type MessageSearch struct {
	Id   int `db:"id"`
	User int
	Room int
}

type RoomSearch struct {
	Id    int `db:"id"`
	Owner int
}
