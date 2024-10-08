package parser

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	filter = "--filter="
	mirror = "--mirror="
	rotate = "--rotate="
	crop   = "--crop="
)

type Option struct {
	Name                                    string
	OffsetX, OffsetY, CropWidth, CropHeight int
	IsHorizontal                            bool
	Rotate                                  int
	Filter                                  string
}

// parser() parses args. Returns n parsed args and a possible error
func Parse(args *[]string) ([]Option, error) {
	var opts []Option
	for _, arg := range (*args)[:len(*args)-2] {
		if strings.HasPrefix(arg, filter) {
			opts = append(opts, Option{Name: "filter", Filter: strings.TrimPrefix(arg, filter)})
		} else if strings.HasPrefix(arg, rotate) {
			arg = strings.TrimPrefix(arg, rotate)
			switch arg {
			case "right", "90", "-270":
				opts = append(opts, Option{Name: "rotate", Rotate: 1})
			case "left", "-90", "270":
				opts = append(opts, Option{Name: "rotate", Rotate: 3})
			case "180", "-180":
				opts = append(opts, Option{Name: "rotate", Rotate: 2})
			default:
				return nil, fmt.Errorf("incorrect value for rotate provided - " + arg)
			}
		} else if strings.HasPrefix(arg, mirror) {
			arg = strings.TrimPrefix(arg, mirror)
			switch arg {
			case "horizontal", "h", "horizontally", "hor":
				opts = append(opts, Option{Name: "mirror", IsHorizontal: true})
			case "vertical", "v", "vertically", "ver":
				opts = append(opts, Option{Name: "mirror", IsHorizontal: false})
			default:
				return nil, fmt.Errorf("incorrect value for mirror provided - " + arg)
			}
		} else if strings.HasPrefix(arg, crop) {
			arg = strings.TrimPrefix(arg, crop)
			values := strings.Split(arg, "-")
			numValues := make([]int, 0)

			if len(values) != 2 && len(values) != 4 {
				// return an error if crop settings are not set properly.
				// it accepts either two or four values
				return nil, fmt.Errorf("incorrect settings for crop provided (2 or 4 is required) - %v", values)
			}
			for _, str := range values {
				num, err := strconv.Atoi(str)
				if err != nil || num < 0 {
					return nil, fmt.Errorf("incorrect setting for crop is provided (non-negative numbers are required) - %v", str)
				}

				numValues = append(numValues, num)
			}
			if len(numValues) < 4 {
				numValues = append(numValues, -1, -1)
			}
			opts = append(opts, Option{Name: "crop", OffsetX: numValues[0], OffsetY: numValues[1], CropWidth: numValues[2], CropHeight: numValues[3]})
		} else {
			return nil, fmt.Errorf("incorrect argument provided for apply - %v", arg)
		}
	}

	return opts, nil
}
