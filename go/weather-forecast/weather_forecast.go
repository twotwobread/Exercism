// Package weather is about weather.
package weather

// CurrentCondition represent weather condition now.
var CurrentCondition string

// CurrentLocation is where you are position.
var CurrentLocation string

// Forecast is function about printing weather condition about your location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
