package main

import (
	"image/color"
	"image/png"
	"os"
)

// try to get message from picture
func seekMessage() ([]string, error) {
	// our res will be stored here
	var resCodes []string

	imgPath := os.Args[2]

	f, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	loadedImg, _ := png.Decode(f)
	size := loadedImg.Bounds().Size()

	// temp var for storing 8-bit code
	var letterCode string

	// loop through pixels
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			// get pixel RGBA value
			pixel := loadedImg.At(x, y)
			originalColor := color.NRGBAModel.Convert(pixel).(color.NRGBA)

			// Get color values
			origValR := intToBin(int64(originalColor.R))
			origValG := intToBin(int64(originalColor.G))
			origValB := intToBin(int64(originalColor.B))
			origValA := intToBin(int64(originalColor.A))
			// Get last bit from color
			codeR := origValR[7]
			codeG := origValG[7]
			codeB := origValB[7]
			codeA := origValA[7]
			// first half of 8-bit code
			codeRGBA := string(codeR) + string(codeG) + string(codeB) + string(codeA)

			// we need to ensure that we have 8-bits for each ASCII code
			// if not => add second half of code
			if len(letterCode) < 8 {
				letterCode += codeRGBA
			} else {
				// if we have 8 bits => append them as element to our res list
				resCodes = append(resCodes, letterCode)
				// letterCode => second half of 8 bits code
				letterCode = codeRGBA
			}
      // if we have last four elements in a row equal to each other => break, 
			// because it means that we've reached end of our message
			if len(resCodes) > 4 {
				if (resCodes[len(resCodes)-1] == resCodes[len(resCodes)-2]) && (resCodes[len(resCodes)-1] == resCodes[len(resCodes)-3]) && (resCodes[len(resCodes)-1] == resCodes[len(resCodes)-4]) {
					return resCodes, nil
				}
			}
		}
	}

	return resCodes, nil
}
