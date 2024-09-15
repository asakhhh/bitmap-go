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
		PrintErrorAndExit(fmt.Sprintf("could not open file %q: ", fileName) + err.Error())
	}

	// Read the BMP Header
	header := make([]byte, 54)
	if _, err := file.Read(header); err != nil {
		PrintErrorAndExit("reading header was not successful: " + err.Error())
	}

	// Check file type for .bmp
	if binary.LittleEndian.Uint16(header[0:2]) != 0x4D42 { // 0x4D42 == "BM"
		PrintErrorAndExit(fmt.Sprintf("%q is not a bitmap file", fileName))
	}

	width := int(binary.LittleEndian.Uint32(header[18:22]))
	height := int(binary.LittleEndian.Uint32(header[22:26]))
	offset := int(binary.LittleEndian.Uint32(header[10:14]))

	compression := int(binary.LittleEndian.Uint32(header[30:34]))
	if compression != 0 {
		PrintErrorAndExit("uncompressed bmp file given. Program accepts only uncompressed 24-bit bmp files.")
	}

	bitsPerPixel := int(binary.LittleEndian.Uint16(header[28:30]))
	if bitsPerPixel != 24 {
		PrintErrorAndExit("given bmp file is not 24-bit. Program accepts only uncompressed 24-bit bmp files.")
	}

	// Calculate row size and read pixel data
	rowSize := ((width*3 + 3) & ^3) // Row size must be a divisible by 4 bytes
	pixelData := make([]byte, rowSize*height)

	// Read the Pixel Data
	file.Seek(int64(offset), 0)
	if _, err := file.Read(pixelData); err != nil {
		PrintErrorAndExit("reading pixel data was not successful: " + err.Error())
	}

	opts, err := parser.Parse(&options)
	if err != nil {
		PrintErrorAndExit(err.Error())
	}

	pixelData = pkg.Mirror(pixelData, width, height, false) // mirror due to reversed row storing in bmp
	for _, opt := range opts {
		switch opt.Name {
		case "filter":
			pixelData, err = pkg.Filter(pixelData, width, height, opt.Filter)
			if err != nil {
				PrintErrorAndExit(err.Error())
			}
		case "mirror":
			pixelData = pkg.Mirror(pixelData, width, height, opt.IsHorizontal)
		case "rotate":
			pixelData, width, height = pkg.Rotate(pixelData, width, height, opt.Rotate)
		case "crop":
			pixelData, width, height, err = pkg.Crop(pixelData, width, height, &opt.OffsetX, &opt.OffsetY, &opt.CropWidth, &opt.CropHeight)
			if err != nil {
				PrintErrorAndExit(err.Error())
			}
		default:
			if strings.HasPrefix(opt.Name, "blur") || strings.HasPrefix(opt.Name, "pixelate") {
				pixelData, err = pkg.Filter(pixelData, width, height, opt.Filter)
				if err != nil {
					PrintErrorAndExit(err.Error())
				}
			} else {
				PrintErrorAndExit("no such filter option - " + opt.Name)
			}
		}
	}
	pixelData = pkg.Mirror(pixelData, width, height, false) // reversing row order back to store correctly afterwards

	binary.LittleEndian.PutUint32(header[18:22], uint32(width))
	binary.LittleEndian.PutUint32(header[22:26], uint32(height))

	file.Close()
	newfile, err := os.Create(options[len(options)-1])
	if err != nil {
		PrintErrorAndExit(fmt.Sprintf("could not create file %q: ", options[len(options)-1]) + err.Error())
	}
	defer newfile.Close()

	_, err = newfile.Write(header)
	if err != nil {
		PrintErrorAndExit(fmt.Sprintf("could not write to file %q: ", options[len(options)-1]) + err.Error())
	}

	empty := make([]byte, offset-54)
	_, err = newfile.Write(empty)
	if err != nil {
		PrintErrorAndExit(fmt.Sprintf("could not write to file %q: ", options[len(options)-1]) + err.Error())
	}

	_, err = newfile.Write(pixelData)
	if err != nil {
		PrintErrorAndExit(fmt.Sprintf("could not write to file %q: ", options[len(options)-1]) + err.Error())
	}
}
