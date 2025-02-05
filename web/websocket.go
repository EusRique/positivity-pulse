package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error connecting WebSocket:", err)
		return
	}
	defer ws.Close()

	// Adiciona o cliente à lista de clientes conectados
	clients[ws] = true

	// Monitorando a conexão do cliente
	for msg := range broadcast {
		if err := ws.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			fmt.Println("Error sending message to client:", err)
			delete(clients, ws) // Remove o cliente da lista em caso de erro
			ws.Close()
			return
		}
	}
}

func NotifyFrontend(msg string) {
	// Envia a mensagem para todos os clientes conectados
	for client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			fmt.Println("Error sending message to client:", err)
			client.Close()
			delete(clients, client) // Remove cliente desconectado ou com erro
		}
	}
}
