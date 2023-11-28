package main

type Room struct {
	name string
	clients map[*Client]bool
	register *Client
	unregister *Client
	broadcast chan Message
}
