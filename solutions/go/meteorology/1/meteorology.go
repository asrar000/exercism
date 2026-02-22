package meteorology

import "fmt"
type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)
func (tU TemperatureUnit) String() string{
    if(tU==0){
        return "°C"
    }
    return "°F"
}
// Add a String method to the TemperatureUnit type

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string{
    value:=fmt.Sprintf("%d %s",t.degree,t.unit.String())
    return value
}
type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (sU SpeedUnit) String() string{
    if(sU==0){
        return "km/h"
    }
    return "mph"
}
type Speed struct {
	magnitude int
	unit      SpeedUnit
}
func (s Speed) String() string {
    value:=fmt.Sprintf("%d %s",s.magnitude,s.unit.String())
    return value
}
// Add a String method to Speed

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}
func (mData MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity",
		mData.location,
		mData.temperature.String(),
		mData.windDirection,
		mData.windSpeed.String(),
		mData.humidity,
	)
}

// Add a String method to MeteorologyData
