package features

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"bitmap/internal/parser"
	"bitmap/pkg"
)

func Apply(options []string) {
	// Open the file
	// fileName := options[len(options)-2]
	fileName := options[len(options)-2]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file %q: %q", fileName, err)
		os.Exit(1)
	}

	// Read the BMP Header
	header := make([]byte, 54)
	if _, err := file.Read(header); err != nil {
		fmt.Println("Error reading header:", err)
		os.Exit(1)
	}

	// Check file type for .bmp
	if binary.LittleEndian.Uint16(header[0:2]) != 0x4D42 { // 0x4D42 == "BM"
		fmt.Printf("Error: %s is not bitmap file\n", fileName)
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

	pixelData = pkg.Mirror(pixelData, width, height, false)

	opts, err := parser.Parse(&options)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, opt := range opts {
		switch opt.Name {
		case "filter":
			pixelData = pkg.Filter(pixelData, width, height, opt.Filter)
		case "mirror":
			pixelData = pkg.Mirror(pixelData, width, height, opt.IsHorizontal)
		case "rotate":
			pixelData, width, height = pkg.Rotate(pixelData, width, height, opt.Rotate)
		case "crop":
			pixelData, width, height, err = pkg.Crop(pixelData, width, height, &opt.OffsetX, &opt.OffsetY, &opt.CropWidth, &opt.CropHeight)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		default:
			if strings.HasPrefix(opt.Name, "blur") || strings.HasPrefix(opt.Name, "pixelate") {
				pixelData = pkg.Filter(pixelData, width, height, opt.Filter)
			} else {
				fmt.Println("No such filter option")
				os.Exit(1)
			}
		}
	}
	pixelData = pkg.Mirror(pixelData, width, height, false)

	binary.LittleEndian.PutUint32(header[18:22], uint32(width))
	binary.LittleEndian.PutUint32(header[22:26], uint32(height))

	file.Close()
	newfile, err := os.Create(options[len(options)-1])
	if err != nil {
		fmt.Printf("Error creating file %q: %q", options[len(options)-1], err)
		os.Exit(1)
	}
	defer newfile.Close()

	_, err = newfile.Write(header)
	if err != nil {
		fmt.Printf("Error writing into file %q: %q", options[len(options)-1], err)
		os.Exit(1)
	}
	_, err = newfile.Write(pixelData)
	if err != nil {
		fmt.Printf("Error writing into file %q: %q", options[len(options)-1], err)
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
