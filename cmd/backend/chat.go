package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MessageResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	RoomID    int       `json:"room_id"`
	CreatedAt time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
	Seen      bool      `json:"seen"`
	Sent      bool      `json:"sent"`
	IsDraft   bool      `json:"is_draft"`
	IsSeen    bool      `json:"is_seen"`
	IsHidden  bool      `json:"is_hidden"`
	IsBlocked bool      `json:"is_blocked"`
	Status    bool      `json:"status"`
	TextBody  string    `json:"text_body"`
}

type RoomResponse struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	MessageCnt int    `json:"msgCnt"`
	LastMsg    string `json:"lastMsg"`
	Members    []int  `json:"members"`
}

func (w *webserver) routesChat(r *gin.Engine) {
	// Grouping routes related to messages
	r.GET("/messages", ginHandler(w.ListMessages))
	r.GET("/messages/:id", ginHandler(w.GetMessageByID))
	r.GET("/messages/user/:id", ginHandler(w.ListMessagesByUserID))
	r.GET("/messages/room/:id", ginHandler(w.ListMessagesByRoomID))

	// Routes for room management
	r.POST("/messages/room/owner/:owner_id/title/:room_name", ginHandler(w.CreateRoom))
	r.POST("/messages/room/:room_id/user/:user_id/p/:p_id", ginHandler(w.InviteMemberToRoom))
	r.DELETE("/messages/room/:room_id/user/:user_id", ginHandler(w.RemoveMemberFromRoom))
	r.GET("/rooms/owner/:owner_id", ginHandler(w.GetRoomsByOwner))
	r.GET("/rooms/user/:user_id", ginHandler(w.GetRoomsByUser))

	// Routes for user management
	// r.POST("/block/user/:user_id/p/:p_id", ginHandler(w.AddBlock))
	// r.DELETE("/block/user/:user_id/p/:p_id", ginHandler(w.RemoveBlock))

	// Versioned API routes
	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("/messages", ginHandler(w.ListMessagesV1))
	// 	v1.GET("/messages/user/:userId", ginHandler(w.GetMessageByIDV1))
	// 	v1.GET("/messages/compose", ginHandler(w.ComposeMessageV1))
	// 	v1.GET("/messages/threads/p/:page/:items/t/:time/:newer", ginHandler(w.ListThreadedMessagesV1))
	// 	v1.GET("/messages/:contact/timestamp", ginHandler(w.GetMessageTimestampV1))
	// 	v1.GET("/messages/:contact/see", ginHandler(w.SeeMessageV1))
	// }

	// Chat-related routes
	r.POST("/chat/send", ginHandler(w.SendMessage))
	r.GET("/ws", func(c *gin.Context) {
		handleWebSocket(c.Writer, c.Request)
	})
}
func (w *webserver) ListMessages(c *gin.Context) error {
	// var resp []MessageResponse
	// params := chatModule.MessageSearch{
	// 	Id: 0,
	// }
	// messages, err := w.message.Find(params)

	// for _, p := range messages {
	// 	resp = append(resp, MessageResponse{
	// 		ID:        p.Id(),
	// 		UserID:    p.User(),
	// 		RoomID:    p.Room(),
	// 		CreatedAt: p.Created(),
	// 		Modified:  p.Modify(),
	// 		Seen:      p.Seen(),
	// 		Sent:      p.Sent(),
	// 		IsDraft:   p.Draft(),
	// 		IsSeen:    p.Seen(),
	// 		IsHidden:  p.Hidden(),
	// 		IsBlocked: p.Blocked(),
	// 		Status:    p.Situation(),
	// 		TextBody:  p.Text(),
	// 	})
	// }

	// if err != nil {
	// 	return err
	// }

	return nil
}
func (w *webserver) GetMessageByID(c *gin.Context) error {
	// id, err := strconv.Atoi(c.Param("id"))
	// params := chatModule.MessageSearch{
	// 	Id: id,
	// }
	// messages, err := w.message.Find(params)
	// if err != nil {
	// 	return err
	// }
	// c.JSON(200, MessageResponse{
	// 	ID:        messages[0].Id(),
	// 	UserID:    messages[0].User(),
	// 	RoomID:    messages[0].Room(),
	// 	CreatedAt: messages[0].Created(),
	// 	Modified:  messages[0].Modify(),
	// 	Seen:      messages[0].Seen(),
	// 	Sent:      messages[0].Sent(),
	// 	IsDraft:   messages[0].Draft(),
	// 	IsSeen:    messages[0].Seen(),
	// 	IsHidden:  messages[0].Hidden(),
	// 	IsBlocked: messages[0].Blocked(),
	// 	Status:    messages[0].Situation(),
	// 	TextBody:  messages[0].Text(),
	// })
	return nil
}
func (w *webserver) ListMessagesByUserID(c *gin.Context) error {
	// var resp []MessageResponse
	// id, err := strconv.Atoi(c.Param("id"))
	// params := chatModule.MessageSearch{
	// 	User: id,
	// }
	// messages, err := w.message.Find(params)
	// if err != nil {
	// 	return err
	// }
	// for _, p := range messages {
	// 	resp = append(resp, MessageResponse{
	// 		ID:        p.Id(),
	// 		UserID:    p.User(),
	// 		RoomID:    p.Room(),
	// 		CreatedAt: p.Created(),
	// 		Modified:  p.Modify(),
	// 		Seen:      p.Seen(),
	// 		Sent:      p.Sent(),
	// 		IsDraft:   p.Draft(),
	// 		IsSeen:    p.Seen(),
	// 		IsHidden:  p.Hidden(),
	// 		IsBlocked: p.Blocked(),
	// 		Status:    p.Situation(),
	// 		TextBody:  p.Text(),
	// 	})
	// }

	return nil
}
func (w *webserver) ListMessagesByRoomID(c *gin.Context) error {
	roomID := c.Param(("roomID"))
	roomIDUint, err := strconv.ParseUint(roomID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return nil
	}
	messages, err := w.chat.GetMessageByRoomID(uint(roomIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	c.JSON(http.StatusOK, messages)
	return nil
}
func (w *webserver) GetRoomsByOwner(c *gin.Context) error {
	// var resp []RoomResponse
	// id, err := strconv.Atoi(c.Param("id"))
	// params := chatModule.RoomSearch{
	// 	Owner: id,
	// }
	// rooms, err := w.room.Find(params)
	// if err != nil {
	// 	return err
	// }
	// for _, p := range rooms {
	// 	resp = append(resp, RoomResponse{
	// 		ID:         p.Id(),
	// 		Name:       p.GetName(),
	// 		MessageCnt: p.GetMessageCnt(),
	// 		LastMsg:    p.GetLastMessage(),
	// 		Members:    p.GetMembers(),
	// 	})
	// }
	return nil
}
func (w *webserver) CreateRoom(c *gin.Context) error {
	var request struct {
		superAdmin int    `uri:"owner_id"`
		name       string `uri:"room_name`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	chatRoom, err := w.chat.CreateRoom(request.superAdmin, request.name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	c.JSON(http.StatusOK, chatRoom)
	return nil
}

func (w *webserver) GetRoomsByUser(c *gin.Context) error {
	// Assuming the user ID is passed as a URL parameter
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return err
	}

	rooms, err := w.chat.GetRoomsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusOK, rooms)
	return nil
}

func (w *webserver) InviteMemberToRoom(c *gin.Context) error {
	pID := c.Param("pID")
	userID := c.Param("userID")
	roomID := c.Param("roomID")
	pIDUint, err := strconv.ParseUint(pID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return nil
	}
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return nil
	}
	roomIDUint, err := strconv.ParseUint(roomID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return nil
	}

	if w.user.CanInvite(uint(pIDUint), uint(userIDUint)) == true {
		err = w.chat.InviteMember(uint(userIDUint), uint(roomIDUint))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return nil
		}
		c.JSON(http.StatusOK, gin.H{"message": "Invitation sent successfully"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "This member has blocked you"})
	}

	return nil
}

func (w *webserver) RemoveMemberFromRoom(c *gin.Context) error {
	userID := c.Param("userID")
	roomID := c.Param("roomID")
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return nil
	}

	roomIDUint, err := strconv.ParseUint(roomID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return nil
	}

	err = w.chat.RemoveMember(uint(userIDUint), uint(roomIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	c.JSON(http.StatusOK, gin.H{"message": "Removed member successfully"})
	return nil
}

func (w *webserver) SendMessage(c *gin.Context) error {
	userID := c.Param("userID")
	roomID := c.Param("roomID")
	content := c.PostForm("content")
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return nil
	}
	roomIDUint, err := strconv.ParseUint(roomID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid room ID"})
		return nil
	}
	if content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message content cannot be empty"})
		return nil
	}
	err = w.chat.SendMessage(uint(userIDUint), uint(roomIDUint), content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
	return nil
}
