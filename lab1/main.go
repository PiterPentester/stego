package main

import (
	"fmt"
)

func main() {
	var message string
	fmt.Println("Enter message to hide:")
	fmt.Scanln(&message)

	lines, err := scanLines("./container.txt")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	// text to ASCII binCodes
	encodedText := textToCode(message)
	fmt.Println(encodedText)
	// ASCII codes to binary
	binCodes := intToBin(encodedText)
	fmt.Println(binCodes)
	// binary to ASCII codes
	decCodes := binToInt(binCodes)
	fmt.Println(decCodes)
	// ASCII codes to text
	decodedText := codeToText(decCodes)
	fmt.Println(decodedText)
}
