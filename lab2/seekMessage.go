package main

import (
	"strings"
)

func seekMessage(text string) string {
	var result string
	ukrLet := "асеіорхАВСЕНІКОРТХ"
	engLet := "aceiopxABCEHIKOPTX"

	// get codes from text
	var codes []string
	var code string

	for _, c := range text {
		char := string(c)
		existInEng := strings.Contains(engLet, char)
		existInUkr := strings.Contains(ukrLet, char)
		if len(code) == 8 {
			if code == "00000000" {
				break
			}
			codes = append(codes, code)
			code = ""
		}
		// if we have an Eng letter that means we have "1"
		if existInEng {
			code += "1"
		} else if existInUkr {
			// if we have Ukr letter and this letter can be replaced by Eng => "0"
			code += "0"
		}
	}
	// binary to ASCII codes
	decCodes := binToInt(codes)

	// ASCII codes to text
	result = codeToText(decCodes)

	return result
}
