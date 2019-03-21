package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give me the DNA sequence you want me to swap")
		os.Exit(1)
	}
	alleles, err := dnaAlleles(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(alleles)
}
