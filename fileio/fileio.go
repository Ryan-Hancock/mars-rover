package fileio

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
)

// ReadFile takes a filename and reads the file and returns a reader for ParseFile
func ReadFile(filename string) (io.Reader, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), err
}

// ParseFile takes a file reader and a separator to return a csv array of strings.
func ParseFile(r io.Reader, seperator rune) ([][]string, error) {
	reader := csv.NewReader(r)
	reader.Comma = seperator
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error ParseTestFile: %s", err.Error())
	}

	return data, err
}
