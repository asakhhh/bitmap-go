package pkg

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Filter(pixelData []byte, width, height int, filter string) []byte {
	rowSize := ((width*3 + 3) & ^3)
	switch filter {
	case "red", "green", "blue":
		col := map[string]int{"red": 2, "green": 1, "blue": 0}
		for y := 0; y < height; y++ {
			for x := 0; x < width*3; x++ {
				if x%3 != col[filter] {
					pixelData[y*rowSize+x] = 0
				}
			}
		}
	case "grayscale":
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				pixel := y*rowSize + x*3
				grayscale := (int(pixelData[pixel]) + int(pixelData[pixel+1]) + int(pixelData[pixel+2])) / 3
				pixelData[pixel] = byte(grayscale)
				pixelData[pixel+1] = byte(grayscale)
				pixelData[pixel+2] = byte(grayscale)
			}
		}
	case "negative":
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				pixel := y*rowSize + x*3
				pixelData[pixel] = byte(255) - pixelData[pixel]
				pixelData[pixel+1] = byte(255) - pixelData[pixel+1]
				pixelData[pixel+2] = byte(255) - pixelData[pixel+2]
			}
		}
	case "pixelate":
		return pixelate(pixelData, width, height, 7)
	case "blur":
		return blur(pixelData, width, height, 5)
	default:
		if strings.HasPrefix(filter, "blur") {
			rad, err := strconv.Atoi(strings.TrimPrefix(filter, "blur"))
			if err != nil || rad < 0 {
				fmt.Println("Invalid blur radius chosen")
				os.Exit(1)
			}
			return blur(pixelData, width, height, rad)
		}
		if strings.HasPrefix(filter, "pixelate") {
			block, err := strconv.Atoi(strings.TrimPrefix(filter, "pixelate"))
			if err != nil || block <= 0 {
				fmt.Println("Invalid block size chosen")
				os.Exit(1)
			}
			return pixelate(pixelData, width, height, block)
		}
	}

	return pixelData
}

func prefixSum(pref *[][][3]int64, y0, x0, y1, x1 int) (int64, int64, int64) {
	return (*pref)[y1+1][x1+1][0] - (*pref)[y1+1][x0][0] - (*pref)[y0][x1+1][0] + (*pref)[y0][x0][0],
		(*pref)[y1+1][x1+1][1] - (*pref)[y1+1][x0][1] - (*pref)[y0][x1+1][1] + (*pref)[y0][x0][1],
		(*pref)[y1+1][x1+1][2] - (*pref)[y1+1][x0][2] - (*pref)[y0][x1+1][2] + (*pref)[y0][x0][2]
}

func blur(pixelData []byte, width, height int, radius int) []byte {
	if radius == 0 {
		return pixelData
	}
	newPixelData := make([]byte, len(pixelData))
	rowSize := ((width*3 + 3) & ^3)

	prefPixel := make([][][3]int64, height+1)
	prefPixel[0] = make([][3]int64, width+1)
	for i := 1; i <= height; i++ {
		prefPixel[i] = make([][3]int64, width+1)
		for col := 0; col < 3; col++ {
			prefPixel[i][0][col] = 0
			for j := 1; j <= width; j++ {
				prefPixel[i][j][col] = prefPixel[i][j-1][col] + int64(pixelData[(i-1)*rowSize+(j-1)*3+col])
			}
		}
	}
	for i := 1; i <= height; i++ {
		for col := 0; col < 3; col++ {
			for j := 0; j <= width; j++ {
				prefPixel[i][j][col] += prefPixel[i-1][j][col]
			}
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := y*rowSize + x*3
			right := min(width-1, x+radius)
			left := max(0, x-radius)
			top := max(0, y-radius)
			bot := min(height-1, y+radius)
			blueSum, greenSum, redSum := prefixSum(&prefPixel, top, left, bot, right)
			cnt := int64(right-left+1) * int64(bot-top+1)
			newPixelData[pixel] = byte(blueSum / cnt)
			newPixelData[pixel+1] = byte(greenSum / cnt)
			newPixelData[pixel+2] = byte(redSum / cnt)
		}
	}
	return newPixelData
}

func pixelate(pixelData []byte, width, height int, block int) []byte {
	blocks := make([][][3]int, height/block+2)
	rowSize := ((width*3 + 3) & ^3)
	for i := range blocks {
		blocks[i] = make([][3]int, width/block+2)
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := y*rowSize + x*3
			blocks[y/block][x/block][0] += int(pixelData[pixel])
			blocks[y/block][x/block][1] += int(pixelData[pixel+1])
			blocks[y/block][x/block][2] += int(pixelData[pixel+2])
		}
	}

	for y := 0; y < (height+block-1)/block; y++ {
		for x := 0; x < (width+block-1)/block; x++ {
			blockHeight := int(math.Min(float64(block), float64(height-y*block)))
			blockWidth := int(math.Min(float64(block), float64(width-x*block)))
			blocksize := blockHeight * blockWidth
			blocks[y][x][0] /= blocksize
			blocks[y][x][1] /= blocksize
			blocks[y][x][2] /= blocksize
		}
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixel := y*rowSize + x*3
			pixelData[pixel] = byte(blocks[y/block][x/block][0])
			pixelData[pixel+1] = byte(blocks[y/block][x/block][1])
			pixelData[pixel+2] = byte(blocks[y/block][x/block][2])
		}
	}
	return pixelData
}
