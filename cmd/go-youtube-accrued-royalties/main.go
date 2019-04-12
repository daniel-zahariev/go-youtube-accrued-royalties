package main

import (
	"fmt"
	royalties "github.com/daniel-zahariev/go-youtube-accrued-royalties"
	"os"
)

var (
	version = "master"
	date    = "1970-01-01"
)

func showVersion() {
	fmt.Printf("go-youtube-accrued-royalties: %v, build %v\n", version, date)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("Process file: ./go-youtube-accrued-royalties <path/to/royalties/file>")
		fmt.Println("Show version: ./go-youtube-accrued-royalties (-v|--version) ")
		return
	} else if (os.Args[1] == "-v") || (os.Args[1] == "--version") {
		showVersion()
		return
	}

	lines, err := royalties.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Read the file '%v' - it has %v lines.\n", os.Args[1], lines)
	}
}
