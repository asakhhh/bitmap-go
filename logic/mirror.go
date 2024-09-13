package logic

func Mirror(pixelData []byte, width, hight int, isHorizontal bool) []byte {
	rowSize := ((width*3 + 3) & ^3) // Row size must be a divisible by 4 bytes
	if isHorizontal {
		return mirrorHorizontal(pixelData, width, hight, rowSize)
	}
	return mirrorVertical(pixelData, width, hight, rowSize)
}

func mirrorHorizontal(pixelData []byte, width, hight, rowSize int) []byte {
	for i := 0; i < hight; i++ {
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

func mirrorVertical(pixelData []byte, width, hight, rowSize int) []byte {
	firstRow := 0
	lastRow := rowSize * (hight - 1)
	for firstRow < lastRow {
		for j := 0; j < width*3; j++ {
			pixelData[firstRow+j], pixelData[lastRow+j] = pixelData[lastRow+j], pixelData[firstRow+j]
		}
		firstRow += rowSize
		lastRow -= rowSize
	}
	return pixelData
}
