package rover

// Min is the Coord for the leftmost side of the plateau
const Min = 0

// Plateau is the area represented by a grid that the rover is able to see.
type Plateau struct {
	X, Y Grid
}

// Grid is the singular square of the grid.
type Grid struct {
	Coord int
}

// NewPlateau initiailses Plateau using the upper rightmost co-ordinates.
func NewPlateau(x, y Grid) Plateau {
	return Plateau{
		X: x,
		Y: y,
	}
}

// GetMax returns the maximum Coord int for X
func (p Grid) GetMax() int {
	return p.Coord
}

// GetMin returns the minimum Coord int for Y
func (p Grid) GetMin() int {
	return Min
}
