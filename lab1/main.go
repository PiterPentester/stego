package main

import (
	"fmt"
	"strconv"
)

// convert our array of integers to string
func codeToText(code []int64) string {
	var result string
	for i := 0; i < len(code); i++ {
		result = result + string(rune(code[i])) // rune(65) == 'A' and so on
	}
	return result
}

// convert our text into array of integers
func textToCode(text string) []int64 {
	runes := []rune(text) // represent our text as runes

	var result []int64

	for i := 0; i < len(runes); i++ {
		result = append(result, int64(runes[i])) // int64('A') == 65
	}

	return result
}

// convert our array of integers into binary
func intToBin(code []int64) string {
	var result string

	for i := 0; i < len(code); i++ {
		result = result + string(strconv.FormatInt(code[i], 2)) // => 77 == '1001101'
	}

	return result
}

//func decode(container string) error {
//	// TODO:
//}
//
func main() {
	message := "My hidden message..."
	fmt.Println(message)

	encodedText := textToCode(message)
	fmt.Println(encodedText)

	binCodes := intToBin(encodedText)
	fmt.Println(binCodes)

	decodedText := codeToText(encodedText)
	fmt.Println(decodedText)
}
