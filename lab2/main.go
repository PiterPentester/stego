package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mode := bufio.NewReader(os.Stdin)
	fmt.Println("Please chose mode. h = hide message, s = seek message:")
	modeAnswer, _ := mode.ReadString('\n')
	modeAnswer = strings.TrimSpace(modeAnswer)

	if modeAnswer == "h" {
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
		data, err := scanContainer("./container.txt")
		if err != nil {
			panic(err)
		}

		// hide message
		res, err := hideMessage(binCodes, data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		writeToFile(res)

	} else if modeAnswer == "s" {
		// scan hidden text
		hidden, err := scanContainer("./hidden.txt")
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
