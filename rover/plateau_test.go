package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPlateau(t *testing.T) {
	assert := assert.New(t)

	topRightX := Grid{Cord: 5}
	topRightY := Grid{Cord: 5}

	plat := NewPlateau(topRightX, topRightY)
	assert.NotNil(plat)
	assert.Equal(plat.X, topRightX)
	assert.Equal(plat.Y, topRightX)
}

func TestGrid_GetMax(t *testing.T) {
	assert := assert.New(t)

	topRightX := Grid{Cord: 6}
	topRightY := Grid{Cord: 5}

	plat := NewPlateau(topRightX, topRightY)
	assert.Equal(plat.X.GetMax(), topRightX.Cord)
}

func TestGrid_GetMin(t *testing.T) {
	assert := assert.New(t)

	topRightX := Grid{Cord: 5}
	topRightY := Grid{Cord: 6}

	plat := NewPlateau(topRightX, topRightY)
	assert.Equal(plat.Y.GetMin(), Min)
	assert.Equal(plat.X.GetMin(), Min)
}
