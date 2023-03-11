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

func RunGoSemanticFuzzTestRoutine(){
	// INIT NUMBER OF ITERATIONS //
	numberOfIteration := 80000

	// START TIMER //
	start := time.Now()

	// LOGGER //
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"GSFTestRoutine:",log.LstdFlags)
	logger.SetOutput(file)

	colorBlue := "\033[34m"
	colorReset := "\033[0m"

	// BEGIN LOGGING //
	helpers.Seperator(1)
	logger.Println("BEGINNING GoSemantic Fuzz TEST ROUTINE")
	helpers.Seperator(1)

	// INIT ASSERTION COUNT
	helpers.InitAssertionsCount()
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",os.Getenv("assertionCount"),"\n",colorReset)
	logger.Println("Number of assertions triggered: ",os.Getenv("assertionCount"))

	// INIT BRANCH COUNT //
	helpers.InitBranchesCount()
	fmt.Print(string(colorBlue),"Number of branches explored: ",os.Getenv("branchCount"),"\n",colorReset)
	logger.Println("Number of branches explored: ",os.Getenv("branchCount"))

	// ACCOUNT EVALUATOR //
	handlers.GoSemanticFuzzAccountGeneratorEvaluator(numberOfIteration)
	helpers.Seperator(1)
	accountEvaluatorAssertions,_ := strconv.Atoi(os.Getenv("assertionCount")) 
	accountEvaluatorBranches,_ := strconv.Atoi(os.Getenv("branchCount"))
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",accountEvaluatorAssertions,"\n",colorReset)
	logger.Println("Number of assertions triggered in account evaluator: ",accountEvaluatorAssertions)
	fmt.Print(string(colorBlue),"Number of branches explored: ",accountEvaluatorBranches+accountEvaluatorAssertions,"\n",colorReset)
	logger.Println("Number of branches explored in account evaluator: ",accountEvaluatorBranches+accountEvaluatorAssertions)
	helpers.Seperator(1)
	time.Sleep(10*time.Second)

	// TRANSACTIONS //
	handlers.GoSemanticFuzzTransactionHandler(numberOfIteration)
	helpers.Seperator(1)
	transactionsAssertions,_ := strconv.Atoi(os.Getenv("assertionCount"))
	transactionsBranches,_ := strconv.Atoi(os.Getenv("branchCount"))
	transactionsAssertions = transactionsAssertions-accountEvaluatorAssertions
	transactionsBranches = transactionsBranches-accountEvaluatorBranches
	fmt.Print(string(colorBlue),"Number of assertions triggered: ",transactionsAssertions,"\n",colorReset)
	logger.Println("Number of assertions triggered in transactions: ",transactionsAssertions)
	fmt.Print(string(colorBlue),"Number of branches explored: ",transactionsAssertions+transactionsBranches,"\n",colorReset)
	logger.Println("Number of branches explored in transactions: ",transactionsAssertions+transactionsBranches)
	helpers.Seperator(1)

	// FINAL SUMMARY //
	helpers.Seperator(1)
	fmt.Println("GoSemanticFuzz Triggered Assertions Summary")
	logger.Println("GoSemanticFuzz Triggered Assertions Summary")
	helpers.Seperator(1)

	fmt.Println("Number of iterations:",numberOfIteration)
	logger.Println("Number of iterations:",numberOfIteration)

	fmt.Print(string(colorBlue),"Number of assertions triggered before starting routine: ",0,"\n",colorReset)
	logger.Println("Number of assertions triggered before starting routine: ",0)

	fmt.Print(string(colorBlue),"Number of branches explored before starting routine: ",0,"\n",colorReset)
	logger.Println("Number of branches explored before starting routine: ",0)

	fmt.Print(string(colorBlue),"Number of assertions triggered in account evaluator: ",accountEvaluatorAssertions,"\n",colorReset)
	logger.Println("Number of assertions triggered in account evaluator: ",accountEvaluatorAssertions)

	fmt.Print(string(colorBlue),"Number of branches explored in account evaluator: ",accountEvaluatorBranches+accountEvaluatorAssertions,"\n",colorReset)
	logger.Println("Number of branches explored in account evaluator: ",accountEvaluatorBranches+accountEvaluatorAssertions)

	fmt.Print(string(colorBlue),"Number of assertions triggered in transactions: ",transactionsAssertions,"\n",colorReset)
	logger.Println("Number of assertions triggered in transactions: ",transactionsAssertions)	
	
	fmt.Print(string(colorBlue),"Number of branches explored in transactions: ",transactionsBranches+transactionsAssertions,"\n",colorReset)
	logger.Println("Number of branches explored in transactions: ",transactionsBranches+transactionsAssertions)

	fmt.Print(string(colorBlue),"Total number of assertions triggered: ",os.Getenv("assertionCount"),"\n",colorReset)
	logger.Println("Total number of assertions triggered: ",os.Getenv("assertionCount"))
	
	totalAssertions,_ := strconv.Atoi(os.Getenv("assertionCount"))
	totalBranches,_ := strconv.Atoi(os.Getenv("branchCount"))
	totalBranches += totalAssertions
	fmt.Print(string(colorBlue),"Total number of branches explored: ",totalBranches,"\n",colorReset)
	logger.Println("Total number of branches explored: ",totalBranches)
	helpers.Seperator(1)

	// END TIMER //
	executionTime := time.Since(start)
	fmt.Println("GSF execution time:",executionTime)
	logger.Println("GSF execution time:",executionTime)
	helpers.Seperator(1)

	// LOG FILE MANAGEMENT
	oldLocation := "result.log"
	newLocation := "logs/goSemanticFuzzResults.log"
	err = os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}

	// ACCOUNTS FILE MANAGEMENT //
	oldLocation = "onlineBanking/storage/accounts.json"
	newLocation = "onlineBanking/storage/goSemanticFuzzAccounts.json"
	err = os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}