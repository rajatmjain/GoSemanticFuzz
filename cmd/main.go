package main

import "GoSemanticFuzz/handlers"

func main(){
	// runGoFuzzTestRoutine()
	// runGoSemanticFuzzTestRoutine()
	handlers.GoSemanticFuzzTransactionHandler(1000)
	//handlers.GoFuzzTransactionHandler(1000)
}

func runGoSemanticFuzzTestRoutine(){
	// ACCOUNT //
	handlers.GoSemanticFuzzAccountGenerator(1000)
	handlers.GoSemanticFuzzAccountGeneratorEvaluator(1000)

	// TRANSACTIONS //
	handlers.GoSemanticFuzzTransactionHandler(10)
}

func runGoFuzzTestRoutine(){
	// ACCOUNT //
	handlers.GoFuzzAccountGenerator(1000)
	handlers.GoFuzzAccountGeneratorEvaluator(1000)

	// TRANSACTIONS //
	handlers.GoFuzzTransactionHandler(10)
}
