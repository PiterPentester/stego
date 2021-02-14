package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mode := bufio.NewReader(os.Stdin)
	fmt.Println("Please chose mode. h = hide message, s = seek message:")
	modeAnswer, _ := mode.ReadString('\n')
	if modeAnswer == "h\n" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter text to hide:")
		message, _ := reader.ReadString('\n')

		// text to ASCII binCodes
		encodedText := textToCode(message)
		//fmt.Println(encodedText)

		// ASCII codes to binary
		binCodes := intToBin(encodedText)
		//fmt.Println(binCodes)

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
		writeToFile(res)

	} else if modeAnswer == "s\n" {
		// scan hidden text
		hidden, err := scanLines("./hidden.txt")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// try to find hidden message
		decodedText := seekMessage(hidden)
		fmt.Println("Message found:")
		fmt.Println(decodedText)
	} else {
		fmt.Println("Error: Choose right mode!")
	}
}
