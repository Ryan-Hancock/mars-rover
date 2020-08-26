package rover

const moveSpace = 1

// Command are the commands sent to the rover to move.
type Command string

// LeftSpin command
const LeftSpin Command = "L"

// RightSpin command
const RightSpin Command = "R"

// Move command
const Move Command = "M"

// Controller type
type Controller struct {
	rover   *Rover
	plateau *Plateau
}

// NewController returns the compass.
func NewController(r Rover, p Plateau) *Controller {
	return &Controller{
		plateau: &p,
		rover:   &r,
	}
}

// Operate decides the action for the command
func (c *Controller) Operate(com Command) error {
	switch com {
	case LeftSpin:
		c.RotateLeft()
	case RightSpin:
		c.RotateRight()
	case Move:
		c.Forward()
	}

	return nil
}

// RotateRight returns the next position for R
func (c *Controller) RotateRight() {
	switch c.rover.direction {
	case North:
		c.rover.SetDirection(East)
	case East:
		c.rover.SetDirection(South)
	case West:
		c.rover.SetDirection(North)
	case South:
		c.rover.SetDirection(West)
	}
}

// RotateLeft returns the next position for L
func (c *Controller) RotateLeft() {
	switch c.rover.direction {
	case North:
		c.rover.SetDirection(West)
	case East:
		c.rover.SetDirection(North)
	case West:
		c.rover.SetDirection(South)
	case South:
		c.rover.SetDirection(East)
	}
}

// Forward returns the cords for the next move for M.
func (c *Controller) Forward() {
	maxX := c.plateau.X.GetMax()
	minX := c.plateau.X.GetMin()
	maxY := c.plateau.Y.GetMax()
	minY := c.plateau.Y.GetMin()
	y := c.rover.y.Cord
	x := c.rover.x.Cord

	// checkPositive checks the current plus next cord agasint the maximum cord
	checkPositive := func(current, max int) bool {
		if (current + moveSpace) > max {
			return false
		}
		return true
	}

	// checkNegative checks the current plus next cord agasint the maximum cord
	checkNegative := func(current, min int) bool {
		if (current - moveSpace) < min {
			return false
		}
		return true
	}

	switch c.rover.direction {
	// Y along the top
	case North:
		if checkPositive(y, maxY) {
			c.rover.SetY(Grid{y + moveSpace})
		}
	case South:
		if checkNegative(y, minY) {
			c.rover.SetY(Grid{y - moveSpace})
		}
	// X along the bottom
	case East:
		if checkPositive(x, maxX) {
			c.rover.SetX(Grid{x + moveSpace})
		}
	case West:
		if checkNegative(x, minX) {
			c.rover.SetX(Grid{x - moveSpace})
		}
	}
}
