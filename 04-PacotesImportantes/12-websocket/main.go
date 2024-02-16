package main

// onbrowser console
// let socket = new WebSocket("ws//localhost:3000/ws")
// socket.onmessage = (event) => { console.log("received from the server:", event.data)}
// socket.send("hello baby from ws clienttt")

// Ref.:
// https://www.youtube.com/watch?v=JuUAEYLkGbM

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New connection from client:", ws.RemoteAddr())
	s.conns[ws] = true
	s.readLoop(ws)
}

func (s *Server) handleWSOrderBook(ws *websocket.Conn) {
	fmt.Println("New connection from client:", ws.RemoteAddr())
	for {
		payload := fmt.Sprintf("orderbook data -> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Microsecond * 500)
	}
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			// connection closed from the client side
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client:", err)
			continue
		}
		msg := buf[:n]

		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		fmt.Println("broadcasting to client")
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Error writing to client:", err)
			}
		}(ws)
	}

}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/wsHandleWSOrderBook", websocket.Handler(server.handleWSOrderBook))
	http.ListenAndServe(":3000", nil)
}
