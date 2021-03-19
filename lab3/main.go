package main

import (
	"bufio"
	"fmt"
	"image/png"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		// show usage help
		defer log.Fatalln("Please check your input")
		log.Println("Usage: ./steGo <mode> path/to/image")
		log.Println("Example: ./steGo hide ./container.png")
		log.Println("Example: ./steGo seek ./container.png")
		log.Println("Now works only with PNG!!!")
	}

	if os.Args[1] == "hide" {
		// get message to hide
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter text to hide:")
		message, _ := reader.ReadString('\n')

		// text to ASCII binCodes
		encodedText := textToCode(message)
		//fmt.Println(encodedText)

		// ASCII codes to binary
		binCodes := codesToBin(encodedText)
		//fmt.Println(binCodes)

		// hide message into data
		img, err := hideMessage(binCodes)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// Create new image
		file, error := os.Create("hidden.png")
		if error != nil {
			log.Fatal(err)
		}
		// Write data to image
		if err := png.Encode(file, img); err != nil {
			file.Close()
			log.Fatal(err)
		}

		if err := file.Close(); err != nil {
			log.Fatal(err)
		}

	} else if os.Args[1] == "seek" {
		// try to find hidden message
		decodedBin, err := seekMessage()
		if err != nil {
			log.Fatal(err)
		}
		// get ASCII codes fom bin list
		decodedASCII := binToInt(decodedBin)
		// get text from ASCII codes
		decodedText := codeToText(decodedASCII)
		fmt.Println("Message found:", decodedText)
	}
}
