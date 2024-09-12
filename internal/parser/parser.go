package parser

import (
	"strings"
)

const (
	filter string = "--filter"
	mirror        = "--mirror"
	rotate        = "--rotate"
	crop          = "--crop"
)

// parser parses args of the apply command
type Parser struct {
	rotate    int // 0 means do not rotate map[int]degrees{0:0, 1:90, 2:180, 3:270}
	horMirror bool
	verMirror bool
	filter    string
	source    string
	dest      string
}

// parser() parses args. Returns n parsed args and an error
func (parser *Parser) parse(args *[]string) (int, error) {
	for i, arg := range *args {
		if strings.HasPrefix(arg, filter) {
			arg = strings.TrimPrefix(arg, filter)
			parser.filter = arg
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
			// in dev
			arg = strings.TrimPrefix(arg, crop)
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
