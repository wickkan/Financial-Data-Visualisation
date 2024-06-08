package websockets

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if string(msg) == "subscribe" {
			// Placeholder: Replace with actual real-time data streaming logic
			for {
				data := map[string]interface{}{
					"symbol": "AAPL",
					"price":  150.0 + float64(len(msg)), // Just a mock data
				}
				if err := conn.WriteJSON(data); err != nil {
					return
				}
			}
		}
	}
}
