package main

import (
	"net/http"

	"github.com/wickkan/Financial-Data-Visualisation/tree/main/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/api/realtime", handlers.RealTimeDataHandler)
	http.HandleFunc("/api/historical", handlers.HistoricalDataHandler)

	http.ListenAndServe(":8080", nil)
}
