package main

import (
	"fmt"
)

var outputChar = "*"

func main() {
	triangle := makeTriangle(checkInput())
	for index := 1; index <= len(triangle); index++ {
		fmt.Println(triangle[index])
	}
}
