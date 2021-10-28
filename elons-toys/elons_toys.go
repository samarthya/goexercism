package elon

import "fmt"

const (
	TotalDistance = "Driven %d meters"
	BatteryCharge = "Battery at %d%%"
)

// 'Drive()' method
func (car *Car) Drive() {
	if car.batteryDrain < car.battery {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	dist := int(car.speed * (car.battery / car.batteryDrain))

	if dist >= trackDistance {
		return true
	}

	return false
}

// 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf(TotalDistance, car.distance)
}

// 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf(BatteryCharge, car.battery)
}
