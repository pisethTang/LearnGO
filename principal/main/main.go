package main

import (
	"fmt"
	// "main/pack"
	"principal/pack"
	// Currently, if we do:
	// "module2/pack2"
	// then we'd get: could not import module2/pack2 (no required module provides package "module2/pack2")
	// that's because we normally can't import a package from one module into another, in this case: from `module2` module into this `principal` module.
	// (one would normally get the same error) if "module2/pack2" turns out to be an external module. 

	// 
)


func main(){
	fmt.Println("In package named --> main")
	pack.F()
	pack2.F()
}