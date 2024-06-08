package main

import (
	"net/http"

	"github.com/wickkan/financial-data-visualisation/handlers"
	"github.com/wickkan/financial-data-visualisation/websockets"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/api/realtime", handlers.RealTimeDataHandler)
	http.HandleFunc("/api/historical", handlers.HistoricalDataHandler)
	http.HandleFunc("/ws", websockets.WebSocketHandler)

	http.ListenAndServe(":8080", nil)
}
