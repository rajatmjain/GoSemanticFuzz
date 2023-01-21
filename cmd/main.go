package main

import "GoSemanticFuzz/handlers"

func main(){
	//transactions.Transfer("tjJ89","gwPu")
	//transactions.Debit("tjJ89")
	//transactions.Credit("tjJ89")
	//handlers.GoFuzzAccountGenerator(100)
	handlers.GoSemanticFuzzAccountGenerator(100)
}
