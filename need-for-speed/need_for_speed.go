package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery, batteryDrain, speed, distance int
}

type CarI interface {
	Drive()
	MaxDistancePossible() int
}

func (c *Car) Drive() {
	if c.batteryDrain > c.battery {
		return
	}
	c.distance += c.speed
	c.battery -= c.batteryDrain
}

func (c Car) MaxDistancePossible() (d int) {
	for i := c.battery; i > 0; i -= c.batteryDrain {
		d += c.speed
	}
	return
}

type Track struct {
	distance int
}

func NewCar(s, b int) Car {
	return Car{speed: s, batteryDrain: b, battery: 100}
}

func NewTrack(d int) Track {
	return Track{distance: d}
}

func Drive(c Car) Car {
	c.Drive()
	return c
}

func CanFinish(c Car, r Track) bool {
	return c.MaxDistancePossible() >= r.distance
}
