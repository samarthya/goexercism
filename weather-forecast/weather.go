// Package weather for the exported vars and function
package weather

var (
	// reflects the current condition.
	CurrentCondition string

	// identifies the current location
	CurrentLocation string
)

// Forecast Forecasts the condition
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
