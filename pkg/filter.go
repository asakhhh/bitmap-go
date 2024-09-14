package pkg

import "fmt"

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
		block := 2
		blocks := make([][][3]int, height/block+2)
		for i := range blocks {
			blocks[i] = make([][3]int, width/block+2)
		}
		// fmt.Println(pixelData)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				pixel := y*rowSize + x*3
				blocks[y/block][x/block][0] += int(pixelData[pixel])
				blocks[y/block][x/block][1] += int(pixelData[pixel+1])
				blocks[y/block][x/block][2] += int(pixelData[pixel+2])
			}
		}

		// fmt.Println(blocks)
		for y := 0; y < (height+block-1)/block; y++ {
			for x := 0; x < (width+block-1)/block; x++ {
				blockHeight := block // int(math.Min(float64(block), float64(height-y+1)))
				blockWidth := block  // int(math.Min(float64(block), float64(width-x+1)))
				blocksize := blockHeight * blockWidth
				blocks[y/block][x/block][0] /= blocksize
				blocks[y/block][x/block][1] /= blocksize
				blocks[y/block][x/block][2] /= blocksize
				// blocks[y/block][x/block][0] /= cnt[y/block][x/block]
				// blocks[y/block][x/block][1] /= cnt[y/block][x/block]
				// blocks[y/block][x/block][2] /= cnt[y/block][x/block]
				fmt.Println(blocks[y/block][x/block])
			}
		}
		// fmt.Println(blocks)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				pixel := y*rowSize + x*3
				pixelData[pixel] = byte(blocks[y/block][x/block][0])
				pixelData[pixel+1] = byte(blocks[y/block][x/block][1])
				pixelData[pixel+2] = byte(blocks[y/block][x/block][2])
			}
		}
	case "blur":
		newPixelData := make([]byte, len(pixelData))
		radius := 10
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				pixel := y*rowSize + x*3
				redSum := 0
				greenSum := 0
				blueSum := 0
				cnt := 0
				for ny := y - radius; ny <= y+radius; ny++ {
					for nx := x - radius; nx <= x+radius; nx++ {
						if ny < 0 || ny >= height || nx < 0 || nx >= width || (x == nx && y == ny) {
							continue
						}
						nPixel := ny*rowSize + nx*3
						cnt++
						blueSum += int(pixelData[nPixel])
						greenSum += int(pixelData[nPixel+1])
						redSum += int(pixelData[nPixel+2])
					}
				}
				newPixelData[pixel] = byte(blueSum / cnt)
				newPixelData[pixel+1] = byte(greenSum / cnt)
				newPixelData[pixel+2] = byte(redSum / cnt)
			}
		}
		return newPixelData
	}

	return pixelData
}
