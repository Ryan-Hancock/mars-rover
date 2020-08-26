package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewController(t *testing.T) {
	assert := assert.New(t)
	r := Rover{x: Grid{5}, y: Grid{5}, direction: North}
	p := Plateau{X: Grid{5}, Y: Grid{5}}
	cont := NewController(r, &p)

	assert.NotNil(cont)
}

func TestController_Operate(t *testing.T) {
	assert := assert.New(t)
	r := Rover{x: Grid{1}, y: Grid{3}, direction: North}
	p := Plateau{X: Grid{5}, Y: Grid{5}}
	cont := NewController(r, &p)

	//testing the rotate left and right functions
	var tests = []struct {
		input    Command
		expected Direction
	}{
		{"L", West},
		{"R", North},
		{"Fail", North},
	}

	for _, test := range tests {
		cont.Operate(test.input)
		assert.Equal(cont.rover.direction, test.expected)
	}

	//testing the move function with y+1 due to being North facing
	curY := cont.rover.y
	cont.Operate(Move)
	assert.Equal(cont.rover.y.Coord, (curY.Coord + 1))
}

func TestController_RotateRight(t *testing.T) {
	assert := assert.New(t)
	r := Rover{x: Grid{1}, y: Grid{3}, direction: North}
	p := Plateau{X: Grid{5}, Y: Grid{5}}
	cont := NewController(r, &p)

	//testing the rotate left and right functions
	var tests = []struct {
		expected Direction
	}{
		{East},
		{South},
		{West},
		{North},
	}

	for _, test := range tests {
		cont.RotateRight()
		assert.Equal(test.expected, cont.rover.direction)
	}
}

func TestController_RotateLeft(t *testing.T) {
	assert := assert.New(t)
	r := Rover{x: Grid{1}, y: Grid{3}, direction: North}
	p := Plateau{X: Grid{5}, Y: Grid{5}}
	cont := NewController(r, &p)

	//testing the rotate left and right functions
	var tests = []struct {
		expected Direction
	}{
		{West},
		{South},
		{East},
		{North},
	}

	for _, test := range tests {
		cont.RotateLeft()
		assert.Equal(test.expected, cont.rover.direction)
	}
}

func TestController_Forward(t *testing.T) {
	assert := assert.New(t)
	r := Rover{x: Grid{3}, y: Grid{3}, direction: North}
	p := Plateau{X: Grid{5}, Y: Grid{5}}
	cont := NewController(r, &p)

	// testing rover near top right corner of a 5,5 grid
	var testsY = []struct {
		expected int
	}{
		{4},
		{5},
		{5},
	}

	for _, test := range testsY {
		cont.Forward()
		assert.Equal(test.expected, cont.rover.y.Coord)
	}

	cont.RotateRight()
	var testsX = []struct {
		expected int
	}{
		{4},
		{5},
		{5},
	}

	for _, test := range testsX {
		cont.Forward()
		assert.Equal(test.expected, cont.rover.x.Coord)
	}
}

func TestController_Forward_Rotate(t *testing.T) {
	assert := assert.New(t)
	r := Rover{x: Grid{3}, y: Grid{3}, direction: North}
	p := Plateau{X: Grid{5}, Y: Grid{5}}
	cont := NewController(r, &p)

	//Moving different directions - currently north 3,3
	var testsRotate = []struct {
		expected int
		input    string
		rotate   func()
	}{
		{4, "x", cont.RotateRight}, //east x4
		{2, "y", cont.RotateRight}, //south y2
		{3, "x", cont.RotateRight}, //west x3
		{3, "y", cont.RotateRight}, //north y3
	}

	for _, test := range testsRotate {
		test.rotate()
		cont.Forward()

		actual := cont.rover.y.Coord
		if test.input == "x" {
			actual = cont.rover.x.Coord
		}

		assert.Equal(test.expected, actual)
	}
}
