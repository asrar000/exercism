// Package weather forecasts the  weather condition at a given location.
package weather 

var (
    // CurrentCondition represents the current weather condition at a given location.
	CurrentCondition string
    // CurrentLocation represents the location of the currently forecasted weather.
	CurrentLocation  string
)
// Forecast returns a string value depicting the current location and its current
// weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
