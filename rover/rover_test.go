package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupMission(t *testing.T) {
	assert := assert.New(t)
	sqd := SetupMission(5, 5)
	s := &Squad{}
	p := Plateau{X: Grid{5}, Y: Grid{5}}
	s.Plataue = p

	assert.Equal(sqd, s)
}

func TestSquad_AddRover(t *testing.T) {
	assert := assert.New(t)
	sqd := SetupMission(5, 5)

	roverX, roverY, roverD := Grid{3}, Grid{2}, North
	_, err := sqd.AddRover(roverX, roverY, roverD)
	assert.Nil(err)
	assert.Len(sqd.Controllers, 1)

	var SouthEast Direction = "SW"
	badX, badY, badD := Grid{6}, Grid{-1}, SouthEast
	_, err = sqd.AddRover(badX, badY, badD)

	assert.NotNil(err)
	assert.Len(sqd.Controllers, 1)
}

func TestSetupRover(t *testing.T) {
	assert := assert.New(t)
	r := SetupRover(Grid{5}, Grid{5}, North)
	assert.Equal(r, Rover{Grid{5}, Grid{5}, North})
}

func TestRover_SetX(t *testing.T) {
	assert := assert.New(t)
	r := SetupRover(Grid{5}, Grid{5}, North)
	r.SetX(Grid{6})

	assert.Equal(r.x.Coord, 6)
}

func TestRover_SetY(t *testing.T) {
	assert := assert.New(t)
	r := SetupRover(Grid{5}, Grid{5}, North)
	r.SetY(Grid{6})

	assert.Equal(r.y.Coord, 6)
}

func TestRover_SetDirection(t *testing.T) {
	assert := assert.New(t)
	r := SetupRover(Grid{5}, Grid{5}, North)
	r.SetDirection(West)

	assert.EqualValues(r, Rover{Grid{5}, Grid{5}, West})
	assert.NotEqualValues(r, Rover{Grid{5}, Grid{5}, North})
}
