package rover

// Min is the Coord for the left side of the max
const Min = 0

// Plateau is the area repersented by a grid the rover can see.
type Plateau struct {
	X, Y Grid
}

// Grid is the singular sqaure of the grid
type Grid struct {
	Coord int
}

// NewPlateau initiailses Plateau using the top right co-ordinates.
func NewPlateau(x, y Grid) Plateau {
	return Plateau{
		X: x,
		Y: y,
	}
}

// GetMax returns the maxaium Coord int for X
func (p Grid) GetMax() int {
	return p.Coord
}

// GetMin returns the maxaium Coord int for Y
func (p Grid) GetMin() int {
	return Min
}
