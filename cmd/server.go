package main

type Server struct {
	rooms map[string]*Room
	wsRegister chan *Client
	wsUnregister chan *Client
}

func createWsServer() *Server {
	return &Server{
		rooms : make(map[string]*Room),
		wsUnregister : make(chan *Client),
		wsRegister : make(chan *Client),
	}
}

func (server *Server) registerClient(client *Client) {
	server.wsRegister <- client
}

func (server *Server) unregisterClient(client *Client) {
	for room := range(client.rooms) {
		room.disconnect(room)
	}
	server.wsUnregister <- client
}

func (server *Server) findRoomByName(roomName string) *Room {
	room, ok := server.rooms[roomName]
	if !ok {
		return nil
	}

	return room
}
