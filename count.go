package main

import (
	"strings"
)


func countWords(s string) int {
	words := strings.Split(s, " ")
	num_words := len(words)
	return num_words
}


// func main(){}


// func main(){
// 	fmt.Printf("Hello world\n")
// }


// 1.4
// error: found packages wordcount (count.go) and main (maina.go) in /home/sething2002/learn-various-tech/LearnGO
// because there can't be 2 packages in one folder.
// to fix it, will have to rename it all to the same name. 

// 1.5
// a main package has 1 or more Go files, all with "package main" in the package clause, and exactly one of them with exactly one main function "func main()". 
// having 2  functions (whether it's main() or countWords()) in 2 different go files in a single main package is not allowed. 
// that's why Go complains. 
// to fix, remove main() from either one of them.

// 1.6
// we can't do go mod init when there's already a go.mod. 
// 

