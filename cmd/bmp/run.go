package bmp

import (
	"os"
	"strings"

	"bitmap/internal/features"
)

func Run() {
	if len(os.Args) == 1 {
		features.PrintHelp("general")
	} else {
		if strings.HasPrefix(os.Args[1], "--help") || os.Args[1] == "-h" {
			features.PrintHelp("general")
		} else {
			for _, arg := range os.Args[2:] {
				if strings.HasPrefix(arg, "--help") || arg == "-h" {
					if os.Args[1] == "apply" || os.Args[1] == "header" {
						features.PrintHelp(os.Args[1])
					} else {
						features.PrintHelp("general")
					}
				}
			}
			if os.Args[1] != "apply" && os.Args[1] != "header" {
				features.PrintErrorAndExit("Incorrect option chosen - " + os.Args[1])
			}
		}
	}
	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "header":
		if len(args) == 0 {
			// no source file provided
			features.PrintErrorAndExit("file not provided")
		}
		features.Header(args[0])
	case "apply":
		if len(args) < 3 {
			features.PrintErrorAndExit("not enough arguments to use this option. Two files and at least one argument are required.")
		}
		features.Apply(args)
	}
}
