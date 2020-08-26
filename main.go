package main

import (
	"fmt"
	"mars-rover/fileio"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing file, < input.txt")
		os.Exit(1)
	}

	r, err := fileio.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("can not read file %s", err.Error())
		os.Exit(1)
	}

	lines, err := fileio.ParseFile(r, rune(' '))
	
	Start(lines)
}
