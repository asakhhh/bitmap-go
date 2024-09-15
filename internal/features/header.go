package features

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
		PrintErrorAndExit(fmt.Sprintf("could not open file %q: ", fileName) + err.Error())
	}
	defer file.Close()

	// Read the BMP Header
	header := make([]byte, bmpHeaderSize)
	if n, err := file.Read(header); err != nil || n < 54 {
		if err != nil {
			PrintErrorAndExit("could not read header")
		} else {
			PrintErrorAndExit("could not read header: incorrect number of bytes")
		}
	}

	// Get Header Data
	fileType := binary.LittleEndian.Uint16(header[0:2])
	if fileType != 0x4D42 { // 0x4D42 == "BM"
		PrintErrorAndExit(fmt.Sprintf("%q is not a bitmap file", fileName))
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
