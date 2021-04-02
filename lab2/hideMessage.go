package main

import (
	"errors"
	"log"
	"os"
	"strings"
)

func hideMessage(binCodes []string, containerText string) (string, error) {
	var result string
	ukrLet := "асеіорхАВСЕНІКОРТХ"
	engLet := "aceiopxABCEHIKOPTX"

	// join or bin codes into one string
	binCodesOneLine := strings.Join(binCodes, "")

	// check that our container are enough to hide message
	if len(binCodesOneLine) > len(containerText) {
		return "", errors.New("container is too small for this message")
	}

	i := 0
	for _, c := range containerText {
		var letter string
		char := string(c)
		exist := strings.Contains(ukrLet, char)
		if exist { // if char in Ukrainian
			if i < len(binCodesOneLine) { // while in rage of binCodesOneLine
				if string(binCodesOneLine[i]) == "1" {
					index := strings.Index(ukrLet, char)
					if index != 0 {
						// for unicode we have 2 indexes for 1 char (in Ukrainian)
						// so we need to divide our index
						index = index / 2
					}
					// set letter equal to similar in Eng
					letter = string(engLet[index])

				} else if string(binCodesOneLine[i]) == "0" {
					// letter equal original language
					letter = char
				}
				i++
			} else {
				// if we have no binCodes => letter == original
				letter = char
			}
		} else {
			// letter also == original if char can't be replaced
			letter = char
		}
		result += letter
	}
	return result, nil
}

func writeToFile(text string) {
	file, err := os.Create("./hidden.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString(text + "\n")
}
