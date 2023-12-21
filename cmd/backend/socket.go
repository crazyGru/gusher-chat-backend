package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

var clients = make(map[*Client]bool)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{Conn: conn, Send: make(chan []byte)}
	clients[client] = true

	go handleMessages(client)
}

func handleMessages(client *Client) {
	defer func() {
		client.Conn.Close()
		delete(clients, client)
	}()

	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("error:%v", err)
			break
		}

		//Broadcast the message to all clients
		for c := range clients {
			select {
			case c.Send <- message:
			default:
				close(c.Send)
				delete(clients, c)
			}
		}
	}
}
