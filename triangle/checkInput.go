package main

import (
	"fmt"
	"os"
	"strconv"
)

var errorMessage = `Argument 1 missing or not an integer value. Please tell us how many lines you want your triangle to be.

Example:
triangle 30`

func checkInput() int {
	if len(os.Args) != 2 {
		fmt.Println(errorMessage)
		os.Exit(1)
	}

	lines, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(errorMessage)
		os.Exit(1)
	}
	return lines
}
