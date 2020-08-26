package rover

// Rover holds the current postions and direction of the mars rover.
type Rover struct {
	x, y      Grid
	direction Direction
}

// SetX sets the x position
func (r *Rover) SetX(x Grid) {
	r.x = x
}

// SetY sets the Y position
func (r *Rover) SetY(y Grid) {
	r.y = y
}

// SetDirection sets the direction pointing
func (r *Rover) SetDirection(direction Direction) {
	r.direction = direction
}
