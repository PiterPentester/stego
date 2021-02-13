package main

func seekMessage(text []string) string {
	var result string

	// get codes from text
	var codes []string
	var code string

	for _, line := range text {
		for i := 0; i < len(line); i++ {
			// check that we can use our search method
			if i+1 < len(line) {
				// each char encoded in 8 bits
				if len(code) == 8 {
					codes = append(codes, code)
					code = ""
				}
				// if char + next char == "  " => we have "1"
				if string(line[i])+string(line[i+1]) == "  " {
					code = code + "1"
					// if char not firts AND char + pervious char != "  " AND char == " " => we have "0"
				} else if (i > 0) && (string(line[i])+string(line[i-1]) != "  ") && string(line[i]) == " " {
					code = code + "0"
				}
			}
		}
	}
	// binary to ASCII codes
	decCodes := binToInt(codes)

	// ASCII codes to text
	result = codeToText(decCodes)

	return result
}
