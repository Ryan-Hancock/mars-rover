package rover

import "fmt"

// Squad holds the Controllers for a group of rovers.
type Squad struct {
	Controllers []*Controller
	Plataue     Plateau
}

// SetupMission sets up the Controllers to command the rovers.
// With the perimeter of the plateau
func SetupMission(xBoundary, yBoundary int) *Squad {
	p := NewPlateau(
		Grid{xBoundary},
		Grid{yBoundary},
	)

	return &Squad{Plataue: p}
}

// AddRover takes the starting location of the rover in x,y and its facing direction.
func (s *Squad) AddRover(xLocation, yLocation Grid, direction Direction) (*Controller, error) {
	if !GetCompass().Directions[direction] {
		return nil, fmt.Errorf("not a support direction %s", direction)
	}

	if xLocation.Coord > s.Plataue.X.GetMax() || xLocation.Coord < s.Plataue.X.GetMin() {
		return nil, fmt.Errorf("rover is outside of plataue on the X asis, %d", xLocation.Coord)
	}

	if yLocation.Coord > s.Plataue.Y.GetMax() || yLocation.Coord < s.Plataue.Y.GetMin() {
		return nil, fmt.Errorf("rover is outside of plataue on the Y asis, %d", xLocation.Coord)
	}

	r := SetupRover(xLocation, yLocation, direction)
	cont := NewController(r, &s.Plataue)

	s.Controllers = append(s.Controllers, cont)
	return cont, nil
}

// Rover holds the current position and direction of the Mars rover.
type Rover struct {
	x, y      Grid
	direction Direction
}

// SetupRover initialises the rover.
func SetupRover(x, y Grid, direction Direction) Rover {
	return Rover{x, y, direction}
}

// SetX sets the x position.
func (r *Rover) SetX(x Grid) {
	r.x = x
}

// SetY sets the Y position.
func (r *Rover) SetY(y Grid) {
	r.y = y
}

// SetDirection sets the facing direction.
func (r *Rover) SetDirection(direction Direction) {
	r.direction = direction
}
