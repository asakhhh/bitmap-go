package io

import (
	"bitmap/logic"
	"fmt"
	"os"
)

func Run() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Help flag")
		return
	}
	if args[0] == "header" {
		if len(args) != 2 {
			fmt.Errorf("Usage: bitmap header <source_file>\nDescription: Prints bitmap file header information")
		}
		logic.Header(args[1])
	} else if args[0] == "apply" {
		if len(args) < 4 {
			fmt.Errorf("Usage: bitmap header <source_file>\nDescription: Prints bitmap file header information")
		}
		logic.Apply(args[1:])
	} else {
		fmt.Println("Error")
		os.Exit(1)
	}
}
