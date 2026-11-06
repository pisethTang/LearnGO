package main

import (
	"fmt"
	// "main/pack"
	"principal/pack"
	// Currently, if we do:
	// "module2/pack2"
	// then we'd get: could not import module2/pack2 (no required module provides package "module2/pack2")
	// that's because we normally can't import a package from one module into another, in this case: from `module2` module into this `principal` module.
	// (one would normally get the same error even if "module2/pack2" turns out to be an external module
	// , it's just in this case it's a local module within the same directory as our principal module).
	// To fix this, one must use a `replace directive` in your go.mod file of the `principal` module, i.e., in /principal/go.mod
	// in my case, I just added: `/home/sething2002/learn-various-tech/LearnGO/module2` into principal/go.mod
	// And then ran "go mod tidy". Here is why if you want to know:
	"module2/pack2"
)


func main(){
	fmt.Println("In package named --> main")
	pack.F()
	pack2.F()
}