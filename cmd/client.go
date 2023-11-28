package main

import "github.com/gorilla/websocket"

type Client struct {
	Name string
	rooms	map[*Room]bool	
	server	*websocket.Conn
}

func newClient(name string, conn *websocket.Conn) *Client {
	return &Client{
		Name: name,
		rooms: make(map[*Room]bool),
		server: conn,
	}
}
