// Package weather tells us current city and its' weather condition.
package weather

// CurrentCondition describes current weather condition.
var CurrentCondition string

// CurrentLocation describes current location to know its' weather condition.
var CurrentLocation string

// Forecast function returns a city name and its' weather condition in a sentence as a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
