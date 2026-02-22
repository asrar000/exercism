package jedlik

import "fmt"

// Drive moves the car once: if there is enough battery, it drains and adds distance.
func (c *Car) Drive() {
	// If there isn't enough battery to drive once, do nothing.
	if c.battery < c.batteryDrain {
		return
	}
	c.battery -= c.batteryDrain
	c.distance += c.speed
}

// DisplayDistance returns the distance driven in meters.
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns the battery percentage.
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish reports whether the car can finish the given track distance.
func (c Car) CanFinish(trackDistance int) bool {
    needToCover:=c.batteryDrain*trackDistance
    tank:=c.battery*c.speed
    return tank-needToCover>=0
	// how many full drives needed to reach or exceed trackDistance
}