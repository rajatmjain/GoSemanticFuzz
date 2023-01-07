package main

import (
	"GoSemanticFuzz/semanticInputGenerator"
	"fmt"
)

func main(){
	fmt.Println(semanticInputGenerator.GenerateFloat64(1.0,2.0,3))
}