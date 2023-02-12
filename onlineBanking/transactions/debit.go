package transactions

import (
	goSemanticFuzz "GoSemanticFuzz/gosemanticfuzz"
	"GoSemanticFuzz/onlineBanking/helpers"
	"GoSemanticFuzz/onlineBanking/misc"
	"fmt"
	"log"
	"os"

	gofuzz "github.com/google/gofuzz"
)

func GoSemanticFuzzDebit(username string){
	// Logger
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"Debit/ASSERTION:",log.LstdFlags)
	logger.SetOutput(file)

	helpers.Seperator(1)
	fmt.Println("DEBIT(GoSemanticFuzz)")
	helpers.Seperator(1)
	account := misc.FetchAccount(username)
	currentAmount := account.Amount
	var debitAmount float64
	float64Schema := goSemanticFuzz.Float64Schema{Minimum: 0,Maximum: currentAmount,Precision: 2}
	float64Fuzzer := goSemanticFuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
	float64Fuzzer.Fuzz(&debitAmount)
	// Assertion
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	if (debitAmount>0 && debitAmount<=currentAmount){
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Valid debit amount","\n",colorReset)
		logger.Println("Valid debit amount")
	}else{
		fmt.Print(string(colorGreen),"Invalid debit amount","\n",colorReset)
		logger.Println("Invalid debit amount")
	}
	newAmount := currentAmount-debitAmount
	account.Amount = newAmount
	fmt.Println("Current amount of",account.Username,":",currentAmount)
	if  err := misc.SaveAccount(account); err==nil{
		fmt.Println(username,"debited with $",debitAmount)
		fmt.Println("New balance for",account.Username,":",account.Amount)
	}
	helpers.Seperator(1)	
}

func GoFuzzDebit(username string){
	// Logger
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"Credit/ASSERTION:",log.LstdFlags)
	logger.SetOutput(file)

	helpers.Seperator(1)
	fmt.Println("DEBIT(GoFuzz)")
	helpers.Seperator(1)
	account := misc.FetchAccount(username)
	currentAmount := account.Amount
	var debitAmount float64
	float64Fuzzer := gofuzz.New()
	float64Fuzzer.Fuzz(&debitAmount)
	// Assertion
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	if (debitAmount>0 && debitAmount<=currentAmount){
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Valid debit amount","\n",colorReset)
		logger.Println("Valid debit amount")
	}else{
		fmt.Print(string(colorGreen),"Invalid debit amount","\n",colorReset)
		logger.Println("Invalid debit amount")
	}
	newAmount := currentAmount-debitAmount
	account.Amount = newAmount
	fmt.Println("Current amount of",account.Username,":",currentAmount)
	if  err := misc.SaveAccount(account); err==nil{
		fmt.Println(username,"debited with $",debitAmount)
		fmt.Println("New balance for",account.Username,":",account.Amount)
	}
	helpers.Seperator(1)	
}