package features

import (
	"fmt"
	"os"
)

func PrintHelp(opt string) {
	if opt == "header" {
		printHeader()
		return
	}
	if opt == "apply" {
		printApply()
		return
	}

	fmt.Println("Usage:")
	fmt.Println("  ./bitmap <option> [args]")
	fmt.Println()
	fmt.Println("The commands are:")
	fmt.Println("  header    prints bitmap file header information.")
	printHeader()
	fmt.Println()
	fmt.Println("  apply     applies processing to the image and saves it to the file.")
	printApply()

	os.Exit(0)
}

func PrintErrorAndExit(msg string) {
	fmt.Println("\u001b[31mError\u001b[0m: " + msg)
	fmt.Printf("Use './bitmap --help' or './bitmap <option> --help' for more information.\n")
	os.Exit(1)
}

func printHeader() {
}

func printApply() {
}
