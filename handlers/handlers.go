package handlers

import (
	"encoding/json"
	"net/http"
)

// Response structure for API responses
type ApiResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Handler for the home endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := ApiResponse{Message: "Welcome to the Financial Data Visualization API"}
	json.NewEncoder(w).Encode(response)
}

// Handler for fetching real-time financial data
func RealTimeDataHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder: Replace with actual logic to fetch real-time data
	data := map[string]interface{}{
		"symbol": "AAPL",
		"price":  150.0,
	}
	response := ApiResponse{Message: "Real-time data fetched successfully", Data: data}
	json.NewEncoder(w).Encode(response)
}

// Handler for fetching historical financial data
func HistoricalDataHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder: Replace with actual logic to fetch historical data
	data := []map[string]interface{}{
		{"date": "2023-06-01", "price": 145.0},
		{"date": "2023-06-02", "price": 148.0},
		{"date": "2023-06-03", "price": 150.0},
	}
	response := ApiResponse{Message: "Historical data fetched successfully", Data: data}
	json.NewEncoder(w).Encode(response)
}
