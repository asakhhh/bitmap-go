package pkg

func Mirror(pixelData []byte, width, height int, isHorizontal bool) []byte {
	rowSize := ((width*3 + 3) & ^3) // Row size must be a divisible by 4 bytes
	if isHorizontal {
		return mirrorHorizontal(pixelData, width, height, rowSize)
	}
	return mirrorVertical(pixelData, width, height, rowSize)
}

func mirrorHorizontal(pixelData []byte, width, height, rowSize int) []byte {
	for i := 0; i < height; i++ {
		start := rowSize * i
		end := start + width*3 - 3
		for start < end {
			pixelData[start], pixelData[end] = pixelData[end], pixelData[start]
			pixelData[start+1], pixelData[end+1] = pixelData[end+1], pixelData[start+1]
			pixelData[start+2], pixelData[end+2] = pixelData[end+2], pixelData[start+2]
			start += 3
			end -= 3
		}
	}
	return pixelData
}

func mirrorVertical(pixelData []byte, width, height, rowSize int) []byte {
	firstRow := 0
	lastRow := rowSize * (height - 1)
	for firstRow < lastRow {
		for j := 0; j < width*3; j++ {
			pixelData[firstRow+j], pixelData[lastRow+j] = pixelData[lastRow+j], pixelData[firstRow+j]
		}
		firstRow += rowSize
		lastRow -= rowSize
	}
	return pixelData
}
