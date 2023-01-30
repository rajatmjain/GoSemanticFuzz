package transactions

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
	"GoSemanticFuzz/onlineBanking/helpers"
	"GoSemanticFuzz/onlineBanking/misc"
	"fmt"

	gofuzz "github.com/google/gofuzz"
)

func GoSemanticFuzzCredit(username string){
	helpers.Seperator(1)
	fmt.Println("CREDIT(GoSemanticFuzz)")
	helpers.Seperator(1)
	var creditAmount float64
	float64Schema := goSemanticFuzz.Float64Schema{Minimum: 0,Maximum: 10000,Precision: 2}
	float64Fuzzer := goSemanticFuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
	float64Fuzzer.Fuzz(&creditAmount)
	// Assertion
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	if (creditAmount>0 && creditAmount<=10000){
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Valid credit amount","\n",string(colorReset))
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Invalid credit amount","\n",string(colorReset))
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
		
}

func GoFuzzCredit(username string){
	helpers.Seperator(1)
	fmt.Println("CREDIT(GoFuzz)")
	helpers.Seperator(1)
	var creditAmount float64
	float64Fuzzer := gofuzz.New()
	float64Fuzzer.Fuzz(&creditAmount)
	// Assertion
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	if (creditAmount>0 && creditAmount<=10000){
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Valid credit amount","\n",string(colorReset))
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Invalid credit amount","\n",string(colorReset))
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
	
}


