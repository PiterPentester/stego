package main

import (
	"fmt"
	"strconv"
)

func codeToText(code []int64) string {
	var result string
	for i := 0; i < len(code); i++ {
		result = result + string(rune(code[i]))
	}
	return result
}

func textToCode(text string) []int64 {
	runes := []rune(text)

	var result []int64

	for i := 0; i < len(runes); i++ {
		result = append(result, int64(runes[i]))
	}

	return result
}

func intToBin(code []int64) string {
	var result string

	for i := 0; i < len(code); i++ {
		result = result + string(strconv.FormatInt(code[i], 2))
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
