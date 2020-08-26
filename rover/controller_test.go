package rover

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewController(t *testing.T) {
	assert := assert.New(t)
	r := Rover{x: Grid{5}, y: Grid{5}, direction: North}
	cont := NewController(r)

	assert.NotNil(cont)
}

func TestController_Operate(t *testing.T) {
	assert := assert.New(t)
	r := Rover{x: Grid{5}, y: Grid{5}, direction: North}
	cont := NewController(r)

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
		fmt.Println(cont.rover)
		assert.Equal(cont.rover.direction, test.expected)
	}
}

func TestController_RotateRight(t *testing.T) {

}

func TestController_RotateLeft(t *testing.T) {

}

func TestController_Forward(t *testing.T) {

}
