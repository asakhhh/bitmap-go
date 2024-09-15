package pkg

func Rotate(pixelData []byte, width, height, rotate int) ([]byte, int, int) {
	if rotate != 1 && rotate != 2 && rotate != 3 {
		return pixelData, width, height
	}
	rowSize := ((width*3 + 3) & ^3)
	newWidth := width
	newHeight := height
	if rotate != 2 {
		newWidth = height
		newHeight = width
	}
	newRowSize := ((newWidth*3 + 3) & ^3)
	newPixelData := make([]byte, newRowSize*newHeight)

	srcIndex := 0
	destIndex := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			srcIndex = y*rowSize + x*3
			if rotate == 3 {
				destIndex = ((newHeight - 1 - x) * newRowSize) + y*3
			} else if rotate == 2 {
				destIndex = ((newHeight - 1 - y) * newRowSize) + ((newWidth - 1 - x) * 3)
			} else if rotate == 1 {
				destIndex = (x * newRowSize) + ((newWidth - 1 - y) * 3)
			}
			newPixelData[destIndex] = pixelData[srcIndex]
			newPixelData[destIndex+1] = pixelData[srcIndex+1]
			newPixelData[destIndex+2] = pixelData[srcIndex+2]
		}
	}
	return newPixelData, newWidth, newHeight
}
