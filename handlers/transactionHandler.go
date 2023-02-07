package handlers

import (
	"GoSemanticFuzz/onlineBanking/authentication"
	"GoSemanticFuzz/onlineBanking/helpers"
	"GoSemanticFuzz/onlineBanking/transactions"
	"GoSemanticFuzz/semanticInputGenerator"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func GoSemanticFuzzTransactionHandler(numberOfTransactions int){
	// LOGGER //
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"GSFTestRoutine:",log.LstdFlags)
	logger.SetOutput(file)

	helpers.Seperator(1)
	fmt.Println("STARTING", numberOfTransactions,"TRANSACTIONS(GoSemanticFuzz)")
	logger.Println("STARTING", numberOfTransactions,"TRANSACTIONS(GoSemanticFuzz)")
	helpers.Seperator(1)
	var i int
	for i=0;i<numberOfTransactions;i++{
		randomTransactionID := semanticInputGenerator.GenerateInt64(1,4)
		switch randomTransactionID{
		case 1:
			goSemanticFuzzCreditHandler()
		case 2:
			goSemanticFuzzDebitHandler()
		case 3:
			goSemanticFuzzTransferHandler()
		default:
			helpers.Seperator(1 )
			fmt.Println("Invalid transaction option")
			logger.Println("Invalid transaction option")
			helpers.Seperator(1)
			colorReset := "\033[0m"
    		colorRed := "\033[31m"
			fmt.Print(string(colorRed),"Total transactions executed successfully: ",i," out of ",numberOfTransactions," transactions",string(colorReset),"\n")
			logger.Println("Total transactions executed successfully: ",i," out of ",numberOfTransactions," transactions")
			helpers.Seperator(1)
			return
		}	
	}
	colorReset := "\033[0m"
    colorRed := "\033[31m"
	fmt.Print(string(colorRed),"Total transactions executed successfully: ",i," out of ",numberOfTransactions," transactions",string(colorReset),"\n")
	logger.Println("Total transactions executed successfully: ",i," out of ",numberOfTransactions," transactions")
	helpers.Seperator(1)
}

func GoFuzzTransactionHandler(numberOfTransactions int){
	// LOGGER //
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"GFTestRoutine:",log.LstdFlags)
	logger.SetOutput(file)

	helpers.Seperator(1)
	fmt.Println("STARTING", numberOfTransactions,"TRANSACTIONS(GoFuzz)")
	logger.Println("STARTING", numberOfTransactions,"TRANSACTIONS(GoFuzz)")
	helpers.Seperator(1)
	var i int
	for i=0;i<numberOfTransactions;i++{
		randomTransactionID := rand.Int31n(4)
		switch randomTransactionID{
		case 1:
			goFuzzCreditHandler()
		case 2:
			goFuzzDebitHandler()
		case 3:
			goFuzzTransferHandler()
		default:
			helpers.Seperator(1)
			fmt.Println("Invalid transaction option")
			logger.Println("Invalid transaction option")
			helpers.Seperator(1)
			colorReset := "\033[0m"
    		colorRed := "\033[31m"
			fmt.Print(string(colorRed),"Total transactions executed successfully: ",i," out of ",numberOfTransactions," transactions",string(colorReset),"\n")
			logger.Println("Total transactions executed successfully: ",i," out of ",numberOfTransactions," transactions")
			helpers.Seperator(1)
			return
		}	
	}
	colorReset := "\033[0m"
    colorRed := "\033[31m"
	fmt.Print(string(colorRed),"Total transactions executed successfully: ",i," out of ",numberOfTransactions," transactions",string(colorReset),"\n")
	logger.Println("Total transactions executed successfully: ",i," out of ",numberOfTransactions," transactions")
	helpers.Seperator(1)
}

func goSemanticFuzzCreditHandler(){
	transactions.GoSemanticFuzzCredit(randomUsername())
}

func goSemanticFuzzDebitHandler(){
	transactions.GoSemanticFuzzDebit(randomUsername())
}

func goSemanticFuzzTransferHandler(){
	transactions.GoSemanticFuzzTransfer(randomUsername(),randomUsername())
}

func goFuzzCreditHandler(){
	transactions.GoFuzzCredit(randomUsername())
}

func goFuzzDebitHandler(){
	transactions.GoFuzzDebit(randomUsername())
}

func goFuzzTransferHandler(){
	transactions.GoFuzzTransfer(randomUsername(),randomUsername())
}

func randomUsername()(string){
	var accounts []authentication.Account
	var usernames []string
	if _, err := os.Stat("onlineBanking/storage/accounts.json"); err == nil {
		accountsFile, _ := os.ReadFile("onlineBanking/storage/accounts.json")
		if err := json.Unmarshal([]byte(accountsFile), &accounts);err!=nil{
			return ""
		}
	}
	for _,acc := range accounts{
		usernames = append(usernames, acc.Username)
	}
	if length := len(usernames);length!=0{
		rando := semanticInputGenerator.GenerateInt64(0,int64(length)+1)
		if rando==0{
			rando = 1
		}
		randomIndex := rand.Int63n(rando)
		return usernames[randomIndex]
	}
	return ""
}