package chat

import (
	"backend/chat/message"
)

const TableNameChat = "chat"

type ChatDto struct {
	ID           int    `gorm:"column:id" json:"id"`
	Owner        int    `gorm:"column:owner" json:"owner"`
	Name         string `gorm:"column:name" json:"name"`
	MessageCnt   int    `gorm:"column:message_cnt" json:"message_cnt"`
	LastMessage  message.Message
	MemberTokens []int `gorm:"column:member_tokens" json:"member_tokens"`
	Notified     bool  `gorm:"column:notified" json:"notified"`
}

func (d *ChatDto) GetID() int                                 { return d.ID }
func (d *ChatDto) SetID(id int)                               { d.ID = id }
func (d *ChatDto) GetOwner() int                              { return d.Owner }
func (d *ChatDto) SetOwner(owner int)                         { d.Owner = owner }
func (d *ChatDto) GetName() string                            { return d.Name }
func (d *ChatDto) SetName(name string)                        { d.Name = name }
func (d *ChatDto) GetMessageCnt() int                         { return d.MessageCnt }
func (d *ChatDto) SetMessageCnt(messageCnt int)               { d.MessageCnt = messageCnt }
func (d *ChatDto) GetLastMessage() message.Message            { return d.LastMessage }
func (d *ChatDto) SetLastMessage(lastMessage message.Message) { d.LastMessage = lastMessage }
func (d *ChatDto) GetMemberTokens() []int                     { return d.MemberTokens }
func (d *ChatDto) SetMemberTokens(memberTokens []int)         { d.MemberTokens = memberTokens }
func (d *ChatDto) GetNotified() bool                          { return d.Notified }
func (d *ChatDto) SetNotified(notified bool)                  { d.Notified = notified }

func (*ChatDto) TableName() string {
	return TableNameChat
}
