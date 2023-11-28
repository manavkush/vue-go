package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func handleWsConnection(wsServer *Server, socket *websocket.Conn) {
	client := newClient(socket)
	registerClient(client)

	go readPump()
	go writePump()
}

func main() {
  flag.Parse()

  http.ListenAndServe(*addr, nil)

	wsServer := createWsServer()

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    socket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error in connection upgrade", err)
			return
		}

		handleWsConnection(wsServer, socket)
  })
}
