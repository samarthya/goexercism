package clock

import ("fmt")

// Representation of Clock.
type Clock struct {
	hour, min int
}

// Depending on the values in the H:M adjust the clock to show the correct time.
// hour <= 24 and min <= 60
func adjustTime(clock *Clock) {
	min := clock.min
	hour := clock.hour

	// If negative we need to adjust the Min and Hours correspondingly.
	for min <0{
		min += 60
		hour -=1
	}

	// If more than 60 means adjust hours and mins to reflect the actual time.
	if min >= 60 {
		hour += (min / 60)
		min %= 60
	}

	for hour < 0 {
		hour += 24
	}

	if hour >= 24 {
		hour %= 24
	} 	
	clock.hour = hour
	clock.min = min
}

// Initialize the clock with the constructor arguments.
func ( clock *Clock) Init(hour, min int) {	
	clock.hour = hour
	clock.min = min
	adjustTime(clock)
}

func ( clock Clock) Add(min int) Clock{
	clock.min += min
	adjustTime(&clock)
	return clock
}

func ( clock Clock) Subtract(min int) Clock{
	clock.min -= min
	adjustTime(&clock)
	return clock
}

// Not returning a pointer as == equality comparator is not going to work.
func New(hour int , min int) Clock {
	clock := new(Clock)
	clock.Init(hour,min)
	return *clock
}

func (clock *Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", clock.hour, clock.min)
}