package main;

var SendMessageAction = "send-message"
var JoinRoomAction = "join-room"
var LeaveRoomAction = "leave-room"

type Message struct {
	Type string	`json:"type"`
	Message string `json:"message"`
	Target string `json:"target"`
	Sender string `json:"sender"`
}
