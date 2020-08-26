package main

import (
	"mars-rover/rover"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	assert := assert.New(t)

	s := rover.SetupMission(5, 5)
	assert.NotNil(s)

	con, err := s.AddRover(rover.Grid{3}, rover.Grid{3}, rover.Direction("N"))
	assert.Nil(err)
	assert.NotNil(con)

	assert.Equal("3 3 N", con.String())
}

func Test_makeRoverAndMove(t *testing.T) {
	assert := assert.New(t)

	s := rover.SetupMission(5, 5)
	con, _ := s.AddRover(rover.Grid{3}, rover.Grid{3}, rover.Direction("N"))
	con.Operate(rover.LeftSpin)
	assert.Equal("3 3 W", con.String())

	con.Operate(rover.Move)
	assert.Equal("2 3 W", con.String())
}

func Test_convertToInt(t *testing.T) {
	assert := assert.New(t)

	i := convertToInt("5")
	assert.Equal(5, i)

	bad := convertToInt("hello")
	assert.Equal(0, bad)
}
