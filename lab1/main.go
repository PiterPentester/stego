package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter text to hide:")
	message, _ := reader.ReadString('\n')

	// text to ASCII binCodes
	encodedText := textToCode(message)
	fmt.Println(encodedText)

	// ASCII codes to binary
	binCodes := intToBin(encodedText)
	fmt.Println(binCodes)

	// scan container
	lines, err := scanLines("./container.txt")
	if err != nil {
		panic(err)
	}

	// hide message
	res, err := hideMessage(binCodes, lines)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(res)

	// binary to ASCII codes
	decCodes := binToInt(binCodes)
	fmt.Println(decCodes)

	// ASCII codes to text
	decodedText := codeToText(decCodes)
	fmt.Println(decodedText)
}
