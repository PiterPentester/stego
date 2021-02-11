package main

import (
	"bufio"
	"os"
)

// scan our container and get text
func scanLines(path string) ([]string, error) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// read by lines
	scanner.Split(bufio.ScanLines)

	var lines []string
	// append lines to array
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
