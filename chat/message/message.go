package message

import "time"

type Message struct {
	ID          int
	SenderID    int
	RoomID      int
	CreatedAt   time.Time
	Attachments Attachment
	TextBody    string
}

type Attachment struct {
	ID       int
	Post     int
	Date     time.Time
	Title    string
	URL      string
	CoverMD  string
	Cover    string
	Type     string
	Hits     int
	Status   string
	Archived bool
	Hidden   bool
}

func (m *Message) Id() int            { return m.ID }
func (m *Message) SenderId() int      { return m.SenderID }
func (m *Message) RoomId() int        { return m.RoomID }
func (m *Message) Created() time.Time { return m.CreatedAt }
func (m *Message) Attach() Attachment { return m.Attachments }
func (m *Message) Text() string       { return m.TextBody }

func FromDto(dto *MessageDto) *Message {
	return &Message{
		ID:        dto.ID,
		SenderID:  dto.SenderID,
		RoomID:    dto.RoomID,
		CreatedAt: dto.Created,
		TextBody:  dto.Text,
	}
}

func ToDto(m *Message) *MessageDto {
	return &MessageDto{
		ID:       m.ID,
		SenderID: m.SenderID,
		RoomID:   m.RoomID,
		Created:  m.CreatedAt,
		Text:     m.TextBody,
	}
}
