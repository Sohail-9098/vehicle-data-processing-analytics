package processor

import (
	"log"

	"github.com/Sohail-9098/vehicle-data-processing-analytics/internal/protobufs/vehicle"
	"github.com/Sohail-9098/vehicle-data-processing-analytics/internal/validator"
)

func ProcessTelemetryData(telemetryData *vehicle.Telemetry) {
	if validator.ValidateTelemetryData(telemetryData) {
		saveToDataTable(telemetryData)
	} else {
		saveToAnamoliesTable(telemetryData)
	}
}

func saveToDataTable(telemetryData *vehicle.Telemetry) {
	log.Println("Valid Data: ", telemetryData)
}

func saveToAnamoliesTable(telemetryData *vehicle.Telemetry) {
	log.Println("Invalid Data: ", telemetryData)
}
