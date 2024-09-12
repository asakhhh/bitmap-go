package logic

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

func Apply(options []string) {
	// Open the file
	fileName := options[len(options)-2]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file %q: %q", fileName, err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the BMP Header
	header := make([]byte, bmpHeaderSize)
	if _, err := file.Read(header); err != nil {
		fmt.Println("Error reading header:", err)
		os.Exit(1)
	}

	width := int(binary.LittleEndian.Uint32(header[18:22]))
	height := int(binary.LittleEndian.Uint32(header[22:26]))
	offset := int(binary.LittleEndian.Uint32(header[10:14]))

	// Calculate row size and read pixel data
	rowSize := ((width*3 + 3) & ^3) // Row size must be a divisible by 4 bytes
	pixelData := make([]byte, rowSize*height)

	// Read the Pixel Data
	file.Seek(int64(offset), 0)
	if _, err := file.Read(pixelData); err != nil {
		fmt.Println("Error reading pixel data:", err)
		os.Exit(1)
	}

	// execute all options
	/*
		for _, option := range options[:len(options)-2] {
			name, value, err := validateOption(option)
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
			switch name {
			case "--mirror":
			case "--filter":
			case "--rotate":
			case "--crop":
			default:
				fmt.Println("invalid option")
				os.Exit(1)
			}
		}
	*/
}

func validateOption(option string) (string, string, error) {
	index := strings.Index(option, "=")
	if index == -1 {
		return "", "", fmt.Errorf("invalid option")
	}
	return option[:index], option[index+1:], nil
}
