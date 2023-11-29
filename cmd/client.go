package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Name string
	rooms	map[*Room]bool	
	socket	*websocket.Conn
	server *Server
	recieve chan []byte
}

func newClient(name string, conn *websocket.Conn, server *Server) *Client {
	return &Client{
		Name: name,
		rooms: make(map[*Room]bool),
		socket: conn,
		server: server,
		recieve: make(chan []byte),
	}
}

// Recieves the message from the readPump in json string format
func (client *Client) handleNewMessage(jsonData []byte) {
	// messageData := <- recieve
	var message *Message

	err := json.Unmarshal(jsonData, message)
	if err != nil {
		log.Printf("Error in unmarshalling json data. %v", err)
		return
	}

	switch message.Type {
		case SendMessageAction:
			client.handleSendMessageAction(message)
		case LeaveRoomAction:
			client.handleLeaveRoomAction(message)
		case JoinRoomAction:
			client.handleJoinRoomAction(message)
	}
}

func (client *Client) handleSendMessageAction(message *Message) {
	roomName := message.Target
	room := client.server.findRoomByName(roomName)

	if room == nil {
		log.Println("Error in handleSendMessageAction. Couldn't find the room.")
		return;
	}

	room.broadcast <- message
}

func (client *Client) handleJoinRoomAction(message *Message) {
	roomName := message.Target
	room := client.server.findRoomByName(roomName)

	if room == nil {
		log.Println("Error in handleJoinRoomAction. Couldn't find the room.")
		return
	}

	client.rooms[room] = true
	room.register <- client
}

func (client *Client) handleLeaveRoomAction(message *Message) {
	roomName := message.Target
	room := client.server.findRoomByName(roomName)

	if room == nil {
		log.Println("Error in handleLeaveRoomAction. Couldn't find the room.")
	}

	present := client.rooms[room]
	if present {
		delete(client.rooms, room)
		room.unregister <- client
	}
}
