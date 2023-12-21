package chat

import (
	"backend/chat/message"
)

type Chat struct {
	id           int
	owner        int
	name         string
	messageCnt   int
	lastMessage  message.Message
	memberTokens []int
	notified     bool
}

func (c *Chat) Id() int                                    { return c.id }
func (c *Chat) Owner() int                                 { return c.owner }
func (c *Chat) Name() string                               { return c.name }
func (c *Chat) MessageCnt() int                            { return c.messageCnt }
func (c *Chat) LastMessage() message.Message               { return c.lastMessage }
func (c *Chat) MemberTokens() []int                        { return c.memberTokens }
func (c *Chat) Notified() bool                             { return c.notified }
func (c *Chat) SetId(id int)                               { c.id = id }
func (c *Chat) SetOwner(owner int)                         { c.owner = owner }
func (c *Chat) SetName(name string)                        { c.name = name }
func (c *Chat) SetMessageCnt(messageCnt int)               { c.messageCnt = messageCnt }
func (c *Chat) SetLastMessage(lastMessage message.Message) { c.lastMessage = lastMessage }
func (c *Chat) SetMemberTokens(memberTokens []int)         { c.memberTokens = memberTokens }
func (c *Chat) SetNotified(notified bool)                  { c.notified = notified }

func (c *Chat) FromDTO(dto *ChatDto) *Chat {
	c.id = dto.ID
	c.owner = dto.Owner
	c.name = dto.Name
	c.messageCnt = dto.MessageCnt
	c.lastMessage = dto.LastMessage
	c.memberTokens = dto.MemberTokens
	c.notified = dto.Notified
	return c
}

func (c *Chat) ToDTO() *ChatDto {
	dto := &ChatDto{
		ID:           c.id,
		Owner:        c.owner,
		Name:         c.name,
		MessageCnt:   c.messageCnt,
		LastMessage:  c.lastMessage,
		MemberTokens: c.memberTokens,
		Notified:     c.notified,
	}
	return dto
}
