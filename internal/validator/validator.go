package validator

import "github.com/Sohail-9098/vehicle-data-processing-analytics/internal/protobufs/vehicle"

const (
	MIN_FUEL_LEVEL  = 10.00
	MAX_SPPED_LEVEL = 300.00
)

func ValidateTelemetryData(telemetryData *vehicle.Telemetry) bool {
	return validateSpeed(telemetryData.Speed) && validateFuelLevel(telemetryData.FuelLevel)
}

func validateSpeed(speed float64) bool {
	return speed <= MAX_SPPED_LEVEL
}

func validateFuelLevel(fuelLevel float64) bool {
	return fuelLevel > MIN_FUEL_LEVEL
}
