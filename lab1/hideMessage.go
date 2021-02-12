package main

import (
	"errors"
	"strings"
)

func hideMessage(binCodes, containerText []string) (string, error) {
	var result string
	zero := " " // 0 => " "
	one := "  " // 1 => "  "

	// split our text to array of words
	var wordsFromContainer []string
	for _, line := range containerText {
		words := strings.Fields(line) // "Test string" => ["Test", "string"]
		for _, v := range words {
			wordsFromContainer = append(wordsFromContainer, v)
		}
	}

	// join or bin codes into one string
	binCodesOneLine := strings.Join(binCodes, "")

	// check that our container are enough to hide message
	if len(binCodesOneLine) > len(wordsFromContainer) {
		return "", errors.New("container is too small for this message")
	}

	var word string
	//  hide message into container by adding " " or "  " to a word
	for i, w := range binCodesOneLine {
		if string(w) == "1" {
			word = wordsFromContainer[i] + one
		} else {
			word = wordsFromContainer[i] + zero
		}
		result = result + word
	}

	return result, nil
}
