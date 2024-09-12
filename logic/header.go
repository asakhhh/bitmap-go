package logic

import (
	"encoding/binary"
	"fmt"
	"os"
)

const (
	bmpHeaderSize = 54
)

func Header(fileName string) {
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the BMP Header
	header := make([]byte, bmpHeaderSize)
	if _, err := file.Read(header); err != nil {
		fmt.Println("Error reading header:", err)
		os.Exit(1)
	}

	// Get Header Data
	fileType := binary.LittleEndian.Uint16(header[0:2])
	if fileType != 0x4D42 { // 0x4D42 == "BM"
		fmt.Printf("Error: %s is not bitmap file\n", fileName)
		os.Exit(1)
	}
	fileSize := int(binary.LittleEndian.Uint32(header[2:6]))
	width := int(binary.LittleEndian.Uint32(header[18:22]))
	height := int(binary.LittleEndian.Uint32(header[22:26]))
	pixelSize := int(binary.LittleEndian.Uint16(header[28:30]))
	imageSize := int(binary.LittleEndian.Uint32(header[34:38]))

	// Output Header Data
	fmt.Println("BMP Header:\n- FileType BM")
	fmt.Println("- FileSizeInBytes", fileSize)
	fmt.Println("- HeaderSize", bmpHeaderSize)
	fmt.Println("DIB Header:\n- DibHeaderSize 40")
	fmt.Println("- WidthInPixels", width)
	fmt.Println("- HeightInPixels", height)
	fmt.Println("- PixelSizeInBits", pixelSize)
	fmt.Println("- ImageSizeInBytes", imageSize)
}
