package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sohail-9098/vehicle-data-processing-analytics/internal/db"
)

func main() {
	http.HandleFunc("/getTelemetry", getTelemetryDataHandler)
	log.Println("server starting on 8080")
	http.ListenAndServe(":8080", nil)
}

func getTelemetryDataHandler(w http.ResponseWriter, r *http.Request) {
	conn := db.NewDB()
	conn.Connect()
	defer conn.Disconnect()
	telemetryData := conn.GetTelemetryData()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(telemetryData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
