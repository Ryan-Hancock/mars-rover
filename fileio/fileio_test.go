package fileio

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	assert := assert.New(t)

	r, err := ReadFile("./test.txt")
	assert.Nil(err)

	b, err := ioutil.ReadAll(r)
	assert.Equal("h", string(b[0]))

	_, err = ReadFile("")
	assert.NotNil(err)
}

func TestParseFile(t *testing.T) {
	assert := assert.New(t)

	data, err := ParseFile(bytes.NewBufferString("hello I am a test file"), rune(' '))
	assert.Nil(err)
	assert.Equal("hello", data[0][0])
	assert.Len(data[0], 6)
}
