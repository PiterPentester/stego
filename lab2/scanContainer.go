package main

import (
	"io/ioutil"
	"strconv"
)

// scan our container and get text
func scanContainer(path string) (string, error) {
	// open file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strconv.Quote(string(data)), nil
}
