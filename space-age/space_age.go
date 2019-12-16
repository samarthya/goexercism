package space

import (
	"math"
)



const earthYears float64 = 31557600

// Planet Type planet defined.
type Planet string

const (
	mercuryYears  float64 = 0.2408467
	venusYears = 0.61519726
	marsYears = 1.8808158
	jupiterYears = 11.862615
	saturnYears = 29.447498
	uranusYears = 84.016846
	neptuneYears = 164.79132
)

const (
	// Mercury planet
	Mercury Planet = "Mercury"

	//Venus planet
	Venus = "Venus"

	//Mars planet
	Mars = "Mars"

	//Saturn planet
	Saturn = "Saturn"

	//Earth planet
	Earth = "Earth"

	//Neptune planet
	Neptune = "Neptune"

	//Uranus planet
	Uranus = "Uranus"

	//Jupiter planet
	Jupiter = "Jupiter"
)

func benchMarkAge( eYears float64, pYear float64) float64 {
	return math.Abs(eYears / pYear)
}

// Age Calculates the age in years.
func Age(age float64, planet Planet) float64 {

	switch planet {

	case Mercury:
		return benchMarkAge( age/earthYears, mercuryYears)
	case Venus:
		return benchMarkAge( age/earthYears, venusYears)
	case Earth:
		return math.Abs(age/earthYears)
	case Mars:
		return benchMarkAge( age/earthYears, marsYears)
	case Jupiter:
		return benchMarkAge( age/earthYears, jupiterYears)
	case Saturn:
		return benchMarkAge( age/earthYears, saturnYears)
	case Uranus:
		return benchMarkAge( age/earthYears, uranusYears)
	case Neptune:
		return benchMarkAge( age/earthYears, neptuneYears)
	}

	return math.Abs(age/earthYears)
}
