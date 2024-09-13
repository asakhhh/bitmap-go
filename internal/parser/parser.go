package parser

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	filter string = "--filter="
	mirror        = "--mirror="
	rotate        = "--rotate="
	crop          = "--crop="
)

// parser parses args of the apply command
type Parser struct {
	rotate     int // 0 means do not rotate map[int]degrees{0:0, 1:90, 2:180, 3:270}
	horMirror  bool
	verMirror  bool
	cropValues []int
	filter     []string
	source     string
	dest       string
}

// parser() parses args. Returns n parsed args and an error
func (parser *Parser) Parse(args *[]string) (int, error) {
	for i, arg := range *args {
		if strings.HasPrefix(arg, filter) {
			arg = strings.TrimPrefix(arg, filter)
			parser.filter = append(parser.filter, arg)
		} else if strings.HasPrefix(arg, rotate) {
			arg = strings.TrimPrefix(arg, rotate)
			switch arg {
			case "right", "90", "-270":
				parser.rotate = (parser.rotate + 1) % 4
			case "left", "-90", "270":
				parser.rotate = (parser.rotate + 3) % 4
			case "180", "-180":
				parser.rotate = (parser.rotate + 2) % 4
			default:
				return i, nil // need to return an error not nil
			}
		} else if strings.HasPrefix(arg, mirror) {
			arg = strings.TrimPrefix(arg, mirror)
			switch arg {
			case "horizontal", "h", "horizontally", "hor":
				parser.horMirror = !parser.horMirror
			case "vertical", "v", "vertically", "ver":
				parser.verMirror = !parser.verMirror
			default:
				return i, nil // need to return an error not nil
			}
		} else if strings.HasPrefix(arg, crop) {
			arg = strings.TrimPrefix(arg, crop)
			values := strings.Split(arg, "-")
			numValues := make([]int, 0)
			fmt.Println(values)

			if len(values) != 2 && len(values) != 4 {
				// return an error if crop settings are not set properly.
				// it accepts either two or four values
				return i, nil // need to return an error not nil
			}

			for _, str := range values {
				num, err := strconv.Atoi(str)
				if err != nil {
					return i, nil // need to return an error not nil
				}

				numValues = append(numValues, num)
			}

			if parser.cropValues == nil {
				parser.cropValues = numValues
			} else {
				if len(numValues) == 2 {
					parser.cropValues[0] += numValues[0]
					parser.cropValues[1] += numValues[1]
				} else {
					if len(parser.cropValues) == 2 {
						parser.cropValues = append(parser.cropValues, numValues[2:4]...)
					} else {
						parser.cropValues[0] += numValues[0]
						parser.cropValues[1] += numValues[1]
						parser.cropValues[2] = numValues[2]
						parser.cropValues[3] = numValues[3]
					}
				}
			}
		} else {
			if i == len(*args)-2 {
				parser.source = arg
			} else if i == len(*args)-1 {
				parser.dest = arg
			} else {
				return i, nil // need to return an error not nil
			}
		}
	}

	return len(*args), nil
}
