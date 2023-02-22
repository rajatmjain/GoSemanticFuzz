package testroutines

import (
	"GoSemanticFuzz/handlers"
	"GoSemanticFuzz/onlineBanking/helpers"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func RunGoFuzzTestRoutine(){
	// START TIMER //
	start := time.Now()

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

	// ACCOUNT EVALUATOR
	handlers.GoFuzzAccountGeneratorEvaluator(10000)
	helpers.Seperator(1)
	accountEvaluatorAssertions,_ := strconv.Atoi(os.Getenv("count")) 
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",accountEvaluatorAssertions,"\n",colorReset)
	logger.Println("Number of assertions triggered in account evaluator: ",accountEvaluatorAssertions)
	helpers.Seperator(1)
	time.Sleep(10*time.Second)

	// TRANSACTIONS //
	handlers.GoFuzzTransactionHandler(10000)
	helpers.Seperator(1)
	transactionsAssertions,_ := strconv.Atoi(os.Getenv("count"))
	transactionsAssertions = transactionsAssertions-accountEvaluatorAssertions
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",transactionsAssertions,"\n",colorReset)
	logger.Println("Number of assertions triggered in transactions: ",transactionsAssertions)
	helpers.Seperator(1)

	// FINAL SUMMARY //
	helpers.Seperator(1)
	fmt.Println("GoFuzz Triggered Assertions Summary")
	logger.Println("GoFuzz Triggered Assertions Summary")
	helpers.Seperator(1)

	fmt.Print(string(colorBlue),"Number of assertions triggered before starting routine: ",0,"\n",colorReset)
	logger.Println("Number of assertions triggered before starting routine: ",0)

	fmt.Print(string(colorBlue),"Number of assertions triggered in account evaluator: ",accountEvaluatorAssertions,"\n",colorReset)
	logger.Println("Number of assertions triggered in account evaluator: ",accountEvaluatorAssertions)

	fmt.Print(string(colorBlue),"Number of assertions triggered in transactions: ",transactionsAssertions,"\n",colorReset)
	logger.Println("Number of assertions triggered in transactions: ",transactionsAssertions)	
	
	fmt.Print(string(colorBlue),"Total number of assertions triggered: ",os.Getenv("count"),"\n",colorReset)
	logger.Println("Total number of assertions triggered: ",os.Getenv("count"))
	helpers.Seperator(1)

	// END TIMER //
	executionTime := time.Since(start)
	fmt.Println("GF execution time:",executionTime)
	logger.Println("GF execution time:",executionTime)
	helpers.Seperator(1)

	// LOG FILE MANAGEMENT //
	oldLocation := "result.log"
	newLocation := "logs/goFuzzResults.log"
	err = os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
	// ACCOUNTS FILE MANAGEMENT //
	oldLocation = "onlineBanking/storage/accounts.json"
	newLocation = "onlineBanking/storage/goFuzzAccounts.json"
	err = os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}