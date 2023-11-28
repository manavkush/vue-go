package main

type Server struct {
	rooms map[*Room]bool
	wsRegister chan *Client
	wsUnregister chan *Client
}

func createWsServer() *Server {
	return &Server{
		room : make(map[*Room]bool),
		wsUnregister : make(chan *Client),
		wsRegister : make(chan *Client),
	}
}

func (server *Server) registerClient(client *Client) {
	server.wsRegister <- client
}

func (server *Server) unregisterClient(client *Client) {
	for room := range(client.rooms) {
		client.disconnect(room)
	}
	server.wsUnregister <- client
}


