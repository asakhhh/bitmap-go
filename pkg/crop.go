package pkg

import "fmt"

func Crop(pixelData []byte, width, height int, offsetX, offsetY, cropWidth, cropHeight *int) ([]byte, int, int, error) {
	// Calculate row size (must be divisible by 4 bytes)
	rowSize := ((width*3 + 3) & ^3)

	// Set default width and height if not provided
	if *cropWidth == -1 {
		temp := width - *offsetX
		cropWidth = &temp
	}
	if *cropHeight == -1 {
		temp := height - *offsetY
		cropHeight = &temp
	}

	if *cropHeight <= 0 || *cropWidth <= 0 {
		return nil, 0, 0, fmt.Errorf("crop width or height are invalid")
	}

	// Check if crop dimensions are valid
	if *offsetX+*cropWidth > width || *offsetY+*cropHeight > height {
		return nil, 0, 0, fmt.Errorf("crop values exceed image size")
	}

	// Create a new array for cropped data
	croppedRowSize := ((*cropWidth * 3) + 3) & ^3 // Ensure cropped row size is divisible by 4
	croppedData := make([]byte, croppedRowSize**cropHeight)

	// Loop through the original pixel data and extract the cropped section
	for i := 0; i < *cropHeight; i++ {
		originalRowStart := (*offsetY + i) * rowSize
		cropRowStart := i * croppedRowSize
		copy(croppedData[cropRowStart:cropRowStart+(*cropWidth*3)], pixelData[originalRowStart+(*offsetX*3):originalRowStart+(*offsetX*3)+(*cropWidth*3)])
	}

	// Return cropped data, new width, new height
	return croppedData, *cropWidth, *cropHeight, nil
}
