package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Println("First argument is:", os.Args[0]) // This is the program name
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
