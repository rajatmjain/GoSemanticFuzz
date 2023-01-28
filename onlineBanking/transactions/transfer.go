package transactions

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
	"GoSemanticFuzz/onlineBanking/misc"
	"fmt"

	gofuzz "github.com/google/gofuzz"
)

func GoSemanticFuzzTransfer(from string, to string){
	misc.Seperator(1)
	fmt.Println("TRANSFER(GoSemanticFuzz)")
	misc.Seperator(1)
	fmt.Println("Transfer initiated from",from,"to",to)
	if(misc.ValidateAccount(from) && misc.ValidateAccount(to)){
		payerAccount := misc.FetchAccount(from)
		payeeAccount := misc.FetchAccount(to)
		payerAmount := payerAccount.Amount
		payeeAmount := payeeAccount.Amount
		fmt.Println("Current amount of", from,":",payerAmount)
		fmt.Println("Current amount of", to,":",payeeAmount)
		var transferAmount float64
		float64Schema := goSemanticFuzz.Float64Schema{Minimum: 0,Maximum: payerAmount,Precision: 2}
		float64Fuzzer := goSemanticFuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
		float64Fuzzer.Fuzz(&transferAmount)
		// Assertion
		if (transferAmount>0 && transferAmount<=payerAmount){
			fmt.Println("Valid transfer amount generated")
		}
		payerAmount -= transferAmount
		payeeAmount += transferAmount
		payerAccount.Amount = payerAmount
		payeeAccount.Amount = payeeAmount
		misc.SaveAccount(payerAccount)
		misc.SaveAccount(payeeAccount)
		fmt.Println("Amount transfered:",transferAmount)
		fmt.Println("New balance for",from,":",payerAmount)
		fmt.Println("New balance for",to,":",payeeAmount)
	}else{
		fmt.Println("Can not validate transaction between accounts.")
	}
	misc.Seperator(1)
	misc.Seperator(1)	
}

func GoFuzzTransfer(from string, to string){
	misc.Seperator(1)
	fmt.Println("TRANSFER(GoFuzz)")
	misc.Seperator(1)
	fmt.Println("Transfer initiated from",from,"to",to)
	if(misc.ValidateAccount(from) && misc.ValidateAccount(to)){
		payerAccount := misc.FetchAccount(from)
		payeeAccount := misc.FetchAccount(to)
		payerAmount := payerAccount.Amount
		payeeAmount := payeeAccount.Amount
		fmt.Println("Current amount of", from,":",payerAmount)
		fmt.Println("Current amount of", to,":",payeeAmount)
		var transferAmount float64
		float64Fuzzer := gofuzz.New()
		float64Fuzzer.Fuzz(&transferAmount)
		// Assertion
		if (transferAmount>0 && transferAmount<=payerAmount){
			fmt.Println("Valid transfer amount generated")
		}
		payerAmount -= transferAmount
		payeeAmount += transferAmount
		payerAccount.Amount = payerAmount
		payeeAccount.Amount = payeeAmount
		misc.SaveAccount(payerAccount)
		misc.SaveAccount(payeeAccount)
		fmt.Println("Amount transfered:",transferAmount)
		fmt.Println("New balance for",from,":",payerAmount)
		fmt.Println("New balance for",to,":",payeeAmount)
	}else{
		fmt.Println("Can not validate transaction between accounts.")
	}
	misc.Seperator(1)
	misc.Seperator(1)	
}