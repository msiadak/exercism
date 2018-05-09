// Package space provides a function that helps calculating ages in years on
// planets in Earth's solar system.
package space

// Planet represents the name of a planet in Earth's solar system.
type Planet string

// EarthPeriod gives the orbital period of Earth in seconds.
const EarthPeriod = 31557600

// PlanetFactors gives the orbital period of each planet relative to Earth.
var PlanetFactors = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age accepts an age in seconds and the name of planet and returns the age in
// years using the orbital period of the planet passed.  Will return 0 if
// an invalid planet name is passed.
func Age(seconds float64, planet Planet) float64 {
	return seconds / (EarthPeriod * PlanetFactors[planet])
}
