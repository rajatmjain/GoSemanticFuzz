package transactions

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
	"GoSemanticFuzz/onlineBanking/misc"
	"fmt"
)

func Debit(username string){
	account := misc.FetchAccount(username)
	currentAmount := account.Amount
	var debitAmount float64
	float64Schema := goSemanticFuzz.Float64Schema{Minimum: 0,Maximum: currentAmount,Precision: 2}
	float64Fuzzer := goSemanticFuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
	float64Fuzzer.Fuzz(&debitAmount)
	newAmount := currentAmount-debitAmount
	account.Amount = newAmount
	if  err := misc.SaveAccount(account); err==nil{
		fmt.Println(username,"debited with $",debitAmount)
		fmt.Println("New amount:",account.Amount)
	}	
}