package main

import (
	"fmt"
	"strings"
)



func toUpperCase(s string) string {
	return strings.ToUpper(s)
}


func main(){
	s := "hello"
	S := toUpperCase(s)
	fmt.Printf("%s --> %s\n", s, S)
}


// Q5
// 1. shout's func main() cannot go in LearnGO/ next to main.go because it would contradict the definition of a package. 
// 2. it needs to go into another folder, i created shout/ in LearnGO/. 
// 3. yes it does. ~/learn-various-tech/LearnGO 
// └─❯ go run ./shout
// hello --> HELLO%  
// 4. 
// - how many modules: 1
// - how many packages: 2
// - how many main packages: 2


// Q5
// not sure? i can just do tree -L 3? 
// 
// 


// ~/learn-various-tech/LearnGO  module & main package
// └─❯ tree -L 2
// .
// ├── count.go
// ├── go.mod
// ├── maina.go
// ├── shout         main package
//      └── upper.go



// Which definition explains why two files in one folder can't both define countWords??
// - shout/ is a module, and its name (from go mod init) is shouter.
// It's also still a main package (D3), because the same folder can be both. LearnGO/ is the same: module wordcount + main package.
// Precise sentence: "shout/ is the module shouter, so by D6 it no longer belongs to the module wordcount."




// Bonus: what would happen if shout/ got its own go.mod? Predict first, then try it, then delete that go.mod.
// - i am in the shouter module. go run .  would work like in LearnGO/ (there's a main package)


// ~/learn-various-tech/LearnGO 
// └─❯ go run ./shout
// main module (wordcount) does not contain package wordcount/shout
// ____ <---- "main" module == current module/module that we are in (wordcount in our case.)