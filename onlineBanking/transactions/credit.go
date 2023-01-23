package transactions

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
	"GoSemanticFuzz/onlineBanking/misc"
	"fmt"
)

func Credit(username string){
	misc.Seperator(1)
	fmt.Println("CREDIT")
	misc.Seperator(1)
	var creditAmount float64
	float64Schema := goSemanticFuzz.Float64Schema{Minimum: 0,Maximum: 10000,Precision: 2}
	float64Fuzzer := goSemanticFuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
	float64Fuzzer.Fuzz(&creditAmount)
	// Assertion
	if (creditAmount>0 && creditAmount<=10000){
		fmt.Println("Valid credit amount")
	}
	account := misc.FetchAccount(username)
	currentAmount := account.Amount
	fmt.Println("Current amount of",account.Username,":",currentAmount)
	newAmount := currentAmount+creditAmount
	account.Amount = newAmount
	if  err := misc.SaveAccount(account); err==nil{
		fmt.Println(username,"credited with $",creditAmount)
		fmt.Println("New balance for",account.Username,":",account.Amount)
	}
	misc.Seperator(1)
	misc.Seperator(1)	
}
