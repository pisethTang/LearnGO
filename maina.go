package main

import (
	"fmt"
)



func main(){
	w := "test string"
	// all_words := strings.Split(w, " ")
	// word_count := len(all_words)
	word_count := countWords(w)
	fmt.Printf("The number of words in %s is %d ---", w, word_count)
}

// 1.1
// The new file is ccalled wordcount, and that name came from the name of the module. 

// 1.2
// D3 breaks because in a given <filename>.go in a folder "f/", and package clause at the top in <filename>.go is "package main"
// we signal to Go that we want to make "f/" the main package (assuming every Go file in the "f/" directory also has "package main")
// and according to D3, a main package has to satisfy 2 criteria:
// - every Go file has "package main" at the top.
// - there exists one and only one Go file with a "func main()" in it. (can't have less than 1 or more than 1). 

// 1.3
// There is no error.
// Changing "package main" to "package wordcount" disqualifies the folder (in this case the root LearnGO/) from being a main package. 
// Moreover, "go build ." searches in the current directory for both main and non-main packages to compile, however we have different cases to deal with:
// - when it's a non-main package, it only compiles, it doesn't create an exectauble out of it. 
// - when it's a main package, it compiles and create an executable.

// 1.4
// 
// 
// 






