package main

import (
	"strconv"
)

// convert our text into array of integers
func textToCode(text string) []int64 {
	runes := []rune(text) // represent our text as runes

	var result []int64

	for i := 0; i < len(runes); i++ {
		result = append(result, int64(runes[i])) // int64('A') == 65
	}

	return result
}

// convert our array of integers into binary array
func intToBin(code []int64) []string {
	var result []string

	for i := 0; i < len(code); i++ {
		result = append(result, strconv.FormatInt(code[i], 2)) // => 77 == '1001101'
	}

	return result
}

// convert our array of binary to integers array
func binToInt(binary []string) []int64 {
	var result []int64

	for i := 0; i < len(binary); i++ {
		n, _ := strconv.ParseInt(binary[i], 2, 64) //=> '1001101' (string) ==> 1001101 (int64)
		num := strconv.FormatInt(n, 10)            // 1001101 (int64) ==> 77 (string)
		n, _ = strconv.ParseInt(num, 10, 64)
		result = append(result, n)
	}

	return result
}

// convert our array of integers to string
func codeToText(code []int64) string {
	var result string
	for i := 0; i < len(code); i++ {
		result = result + string(rune(code[i])) // rune(65) == 'A' and so on
	}
	return result
}
