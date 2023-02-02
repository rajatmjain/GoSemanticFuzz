package testroutines

import (
	"GoSemanticFuzz/handlers"
	"GoSemanticFuzz/onlineBanking/helpers"
	"fmt"
	"log"
	"os"
	"time"
)

func RunGoFuzzTestRoutine(){
	// LOGGER //
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"GFTestRoutine:",log.LstdFlags)
	logger.SetOutput(file)

	colorBlue := "\033[34m"
	colorReset := "\033[0m"

	// BEGIN LOGGING //
	helpers.Seperator(1)
	logger.Println("BEGINNING GoFuzz TEST ROUTINE")
	helpers.Seperator(1)

	// INIT ASSERTION COUNT
	helpers.InitAssertionsCount()
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",os.Getenv("count"),"\n",colorReset)
	logger.Println("Number of assertions triggered: ",os.Getenv("count"))

	// ACCOUNT //
	handlers.GoFuzzAccountGenerator(1000)
	helpers.Seperator(1)
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",os.Getenv("count"),"\n",colorReset)
	logger.Println("Number of assertions triggered: ",os.Getenv("count"))
	helpers.Seperator(1)
	time.Sleep(10*time.Second)

	// ACCOUNT EVALUATOR
	handlers.GoFuzzAccountGeneratorEvaluator(1000)
	helpers.Seperator(1)
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",os.Getenv("count"),"\n",colorReset)
	logger.Println("Number of assertions triggered: ",os.Getenv("count"))
	helpers.Seperator(1)

	// TRANSACTIONS //
	time.Sleep(10*time.Second)
	handlers.GoFuzzTransactionHandler(10)
	helpers.Seperator(1)
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",os.Getenv("count"),"\n",colorReset)
	logger.Println("Number of assertions triggered: ",os.Getenv("count"))
	helpers.Seperator(1)
}