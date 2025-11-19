// package echo3
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//  a more efficient version of echo2
	// fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Name of program %s\n", os.Args[0])
	// fmt.Println(os.Args[1:])

	// Exercise 2: printing the index and value of each cmd line arguments
	// for index, arg := range os.Args{
	// 	fmt.Printf("Arg %d: %s\n", index, arg)
	// }

	// Exercise 3: measuring execution time between 
	// loop based and strings.Join based implementations
	fmt.Println("Measuring execution time of two implementations:")
	start := time.Now()
	s, sep := "", " "
	for _, arg := range os.Args[1:]{
		s += sep + arg
		// sep = " "
	}
	fmt.Println(s)
	elapsed := time.Since(start)
	
	
	
	start2 := time.Now()
	fmt.Println("Using strings.Join:", os.Args[1:])
	elapsed2 := time.Since(start2)


	// print elapsed times
	fmt.Printf("Loop based implementation took %s\n", elapsed)
	fmt.Printf("strings.Join based implementation took %s\n", elapsed2)


	// run on bigger input
	// go run echo3.go $(seq 1 10000)
}
