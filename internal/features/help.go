package features

import (
	"fmt"
	"os"
)

func PrintHelp(opt string) {
	if opt == "header" {
		printHeader()
	} else if opt == "apply" {
		printApply()
	} else {
		fmt.Println("Usage:")
		fmt.Println("  ./bitmap <option> [args]")
		fmt.Println("\nThe commands are:")
		fmt.Println("  header    prints bitmap file header information.")
		fmt.Println("  apply     applies processing to the image and saves it to the file.")
		fmt.Println("Use './bitmap <option> --help' for more info about a specific option.")
	}
	os.Exit(0)
}

func PrintErrorAndExit(msg string) {
	fmt.Println("\u001b[31mError\u001b[0m: " + msg)
	fmt.Printf("Use './bitmap --help' or './bitmap <option> --help' for more information.\n")
	os.Exit(1)
}

func printHeader() {
	fmt.Println("Usage:\n  ./bitmap header <source_file>\t- outputs header info of the file provided\n  ./bitmap header --help\t- outputs help info for the header option")
}

func printApply() {
	fmt.Println("Usage:\n  ./bitmap apply [args] <src> <dest> - applies specified transformations to the <src> and writes the result in <dest>")
	fmt.Println("  ./bitmap apply --help\t\t     - outputs help info for the apply option")
	fmt.Println("\nTransformations:")
	fmt.Println("\t--crop=\u001b[34mOX-OY[-W-H]\u001b[0m - crops the photo based on the offset \u001b[34m(OX, OY)\u001b[0m [and width/height \u001b[34m(W, H)\u001b[0m if provided].")
	fmt.Println("\t--filter=\u001b[34m<opt>\u001b[0m     - applies a filter.   \u001b[34m<opt>\u001b[0m=\u001b[36m[ blue | red | green | grayscale | negative | pixelate[N] | blur[N] ]\u001b[0m")
	fmt.Println("\t--mirror=\u001b[34m<opt>\u001b[0m     - mirrors the bitmap. \u001b[34m<opt>\u001b[0m=\u001b[36m[ h | hor | horizontal | horizontally | v | ver | vertical | vertically ]\u001b[0m")
	fmt.Println("\t--rotate=\u001b[34m<opt>\u001b[0m     - rotates the bitmap. \u001b[34m<opt>\u001b[0m=\u001b[36m[ right | 90 | 180 | 270 | left | -90 | -180 | -270 ]\u001b[0m")
}

func PrintWarning(msg string) {
	fmt.Print("\x1b[35mWarning\u001b[0m: " + msg)
}
