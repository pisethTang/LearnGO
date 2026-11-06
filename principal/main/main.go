package main

import (
	"fmt"
	"main/pack"
	// "principal/pack"


	// Currently, if we do:
	// "module2/pack2" 
	// then we'd get: could not import module2/pack2 (no required module provides package "module2/pack2")
	
)


func main(){
	fmt.Println("In package named --> main")
	pack.F()
	pack2.F()
}