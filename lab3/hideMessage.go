package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

// hide data in PNG
func hideMessage(binCodes []string) (image.Image, error) {
	// list of bin codes become one long string
	var codesString string
	for _, elem := range binCodes {
		codesString += elem
	}

	imgPath := os.Args[2]

	f, err := os.Open(imgPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// decode img
	loadedImg, _ := png.Decode(f)
	size := loadedImg.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	img := image.NewNRGBA(rect)
	// Ensure that we have enough space to our message
	if len(codesString) < size.Y*size.X {
		for y := 0; y < size.Y; y++ {
			for x := 0; x < size.X; x++ {
				// get pixel coordinates
				pixel := loadedImg.At(x, y)
				// get pixel colors
				originalColor := color.NRGBAModel.Convert(pixel).(color.NRGBA)
				// assign first four elements from our bin string to R, G, B, A
				if len(codesString) >= 4 {
					R := codesString[0]
					G := codesString[1]
					B := codesString[2]
					A := codesString[3]
					// remove used elements from our bin string
					if len(codesString) != 4 {
						codesString = codesString[4:]
						// get orignal color values
						origValR := intToBin(int64(originalColor.R))
						origValG := intToBin(int64(originalColor.G))
						origValB := intToBin(int64(originalColor.B))
						origValA := intToBin(int64(originalColor.A))
						// assign new data to last bits
						origValR = origValR[:7] + string(R)
						origValG = origValG[:7] + string(G)
						origValB = origValB[:7] + string(B)
						origValA = origValA[:7] + string(A)
						// Convert data types
						newValR, _ := strconv.ParseInt(origValR, 2, 64)
						newValG, _ := strconv.ParseInt(origValG, 2, 64)
						newValB, _ := strconv.ParseInt(origValB, 2, 64)
						newValA, _ := strconv.ParseInt(origValA, 2, 64)
						// Write modified data
						img.Set(x, y, color.NRGBA{
							R: uint8(newValR),
							G: uint8(newValG),
							B: uint8(newValB),
							A: uint8(newValA),
						})
					} else {
						// write unmodified data
						img.Set(x, y, color.NRGBA{
							R: uint8(originalColor.R),
							G: uint8(originalColor.G),
							B: uint8(originalColor.B),
							A: uint8(originalColor.A),
						})
					}
				}
			}
		}
	}

	return img, nil
}
