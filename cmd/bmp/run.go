package bmp

import (
	"bitmap/logic"
	"fmt"
	"os"
)

func Run() {
	if len(os.Args) <= 2 {
		fmt.Println("Help flag")
		return
	}
	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "header":
		if len(args) == 0 {
			// no source file provided
			fmt.Printf("Usage: bitmap header <source_file>\nDescription: Prints bitmap file header information\n")
			os.Exit(1)
		}
		logic.Header(args[0])
	// in dev
	case "apply":
		if len(args) < 3 {
			fmt.Printf("Usage: bitmap apply [--rotate-right|--rotate-left] <input_file> <output_file>\nDescription: Applies transformations to the bitmap file\n")
			os.Exit(1)
		}
		logic.Apply(args[1:])
	// in dev
	default:
		fmt.Println("Error: Unknown command")
		fmt.Println("Usage: bitmap <command> [options]")
		os.Exit(1)
	}
}
