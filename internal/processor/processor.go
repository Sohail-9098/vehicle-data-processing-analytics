package processor

import (
	"github.com/Sohail-9098/vehicle-data-processing-analytics/internal/db"
	"github.com/Sohail-9098/vehicle-data-processing-analytics/internal/protobufs/vehicle"
	"github.com/Sohail-9098/vehicle-data-processing-analytics/internal/validator"
)

func ProcessTelemetryData(telemetryData *vehicle.Telemetry) {
	conn := db.NewDB()
	conn.Connect()
	defer conn.Disconnect()

	if validator.ValidateTelemetryData(telemetryData) {
		saveToDataTable(conn, telemetryData)
	} else {
		saveToAnamoliesTable(conn, telemetryData)
	}
}

func saveToDataTable(conn *db.DB, telemetryData *vehicle.Telemetry) {
	conn.InsertTelemetryData("TELEMETRY", telemetryData)
}

func saveToAnamoliesTable(conn *db.DB, telemetryData *vehicle.Telemetry) {
	conn.InsertTelemetryData("ANAMOLIES", telemetryData)
}
