package speed

// TODO: define the 'Car' type struct
type Car struct{
        speed int 
        batteryDrain int
        battery int
        distance int
    }

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
    /*
    type Car struct{
        speed int 
        batteryDrain int
        battery int
        distance int
    }
    */
    return Car{speed,batteryDrain,100,0}
    
	panic("Please implement the NewCar function")
}

// TODO: define the 'Track' type struct
type Track struct{
        distance int
    }
    
// NewTrack creates a new track
func NewTrack(distance int) Track {
    
    
    return Track{distance}
	panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if(car.batteryDrain<=car.battery){
        car.battery=car.battery-car.batteryDrain
        car.distance=car.distance+car.speed
    }
    return car
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    numberOfTimes:=float64(car.battery)/float64(car.batteryDrain)
    return numberOfTimes*float64(car.speed)>=float64(track.distance)
	panic("Please implement the CanFinish function")
}
