package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCompass(t *testing.T) {
	assert := assert.New(t)

	comp := GetCompass()
	assert.NotNil(comp)
}

func TestCompass(t *testing.T) {
	assert := assert.New(t)

	comp := GetCompass()

	assert.True(comp.Directions[North])
	assert.True(comp.Directions[West])
	assert.True(comp.Directions[South])
	assert.True(comp.Directions[East])

	assert.False(comp.Directions["T"])
}
