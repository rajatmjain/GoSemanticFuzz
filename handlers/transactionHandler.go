package handlers

import (
	"GoSemanticFuzz/onlineBanking/authentication"
	"GoSemanticFuzz/onlineBanking/misc"
	"GoSemanticFuzz/onlineBanking/transactions"
	"GoSemanticFuzz/semanticInputGenerator"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

func TransactionHandler(numberOfTransactions int){
	misc.Seperator(1)
	fmt.Println("STARTING", numberOfTransactions,"TRANSACTIONS")
	misc.Seperator(1)
	for i:=0;i<numberOfTransactions;i++{
		randomTransactionID := semanticInputGenerator.GenerateInt64(1,4)
		switch randomTransactionID{
		case 1:
			creditHandler()
		case 2:
			debitHandler()
		case 3:
			transferHandler()
		default:
			return
		}	
	}
}

func creditHandler(){
	transactions.Credit(randomUsername())
}

func debitHandler(){
	transactions.Debit(randomUsername())
}

func transferHandler(){
	transactions.Transfer(randomUsername(),randomUsername())
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
		randomIndex := rand.Int63n(semanticInputGenerator.GenerateInt64(0,int64(length)+1))
		return usernames[randomIndex]
	}
	return ""
}