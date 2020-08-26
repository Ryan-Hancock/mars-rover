package rover

import (
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

	
}

func TestController_RotateRight(t *testing.T) {

}

func TestController_RotateLeft(t *testing.T) {

}

func TestController_Forward(t *testing.T) {

}
