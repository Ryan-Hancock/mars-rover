package rover

// North ...
const North Direction = "N"

// East ...
const East Direction = "E"

// West ...
const West Direction = "W"

// South ...
const South Direction = "S"

// Direction is the letter for the postion on a compass
type Direction string

// Compass holds the directions represented by strings.
type Compass struct {
	Directions map[Direction]bool
}

// directions is a map Compass points, mapped to bools for safe checking.
var directions = map[Direction]bool{
	North: true,
	East:  true,
	West:  true,
	South: true,
}

// GetCompass returns the compass.
func GetCompass() Compass {
	return Compass{Directions: directions}
}
