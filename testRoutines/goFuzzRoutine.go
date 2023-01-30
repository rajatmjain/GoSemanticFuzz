package testroutines

import (
	"GoSemanticFuzz/handlers"
	"GoSemanticFuzz/onlineBanking/helpers"
	"fmt"
	"os"
	"time"
)

func RunGoFuzzTestRoutine(){
	colorBlue := "\033[34m"
	colorReset := "\033[0m"

	// INIT ASSERTION COUNT
	helpers.InitAssertionsCount()
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",os.Getenv("count"),"\n",colorReset)

	// ACCOUNT //
	handlers.GoFuzzAccountGenerator(1000)
	helpers.Seperator(1)
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",os.Getenv("count"),"\n",colorReset)
	helpers.Seperator(1)
	time.Sleep(10*time.Second)

	// ACCOUNT EVALUATOR
	handlers.GoFuzzAccountGeneratorEvaluator(1000)
	helpers.Seperator(1)
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",os.Getenv("count"),"\n",colorReset)
	helpers.Seperator(1)

	// TRANSACTIONS //
	time.Sleep(10*time.Second)
	handlers.GoFuzzTransactionHandler(10)
	helpers.Seperator(1)
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",os.Getenv("count"),"\n",colorReset)
	helpers.Seperator(1)
}