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

func GoSemanticFuzzTransfer(from string, to string){
	// Logger
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"Transfer/ASSERTION:",log.LstdFlags)
	logger.SetOutput(file)
	
	helpers.Seperator(1)
	fmt.Println("TRANSFER(GoSemanticFuzz)")
	helpers.Seperator(1)
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
		colorGreen := "\033[32m"
		colorReset := "\033[0m"
		if (transferAmount>0 && transferAmount<=payerAmount){
			helpers.CountAssertions()
			fmt.Print(string(colorGreen),"Valid transfer amount","\n",string(colorReset))
			logger.Println("Valid transfer amount")
			logger.Println("Transfer amount:",transferAmount)
		}else{
			fmt.Print(string(colorGreen),"Invalid transfer amount","\n",string(colorReset))
			logger.Println("Invalid transfer amount")
			logger.Println("Transfer amount:",transferAmount)
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
	helpers.Seperator(1)	
}

func GoFuzzTransfer(from string, to string){
	// Logger
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"Transfer/ASSERTION:",log.LstdFlags)
	logger.SetOutput(file)

	helpers.Seperator(1)
	fmt.Println("TRANSFER(GoFuzz)")
	helpers.Seperator(1)
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
		colorGreen := "\033[32m"
		colorReset := "\033[0m"
		if (transferAmount>0 && transferAmount<=payerAmount){
			helpers.CountAssertions()
			fmt.Print(string(colorGreen),"Valid transfer amount","\n",string(colorReset))
			logger.Println("Valid transfer amount")
			logger.Println("Transfer amount:",transferAmount)
		}else{
			fmt.Print(string(colorGreen),"Invalid transfer amount","\n",string(colorReset))
			logger.Println("Invalid transfer amount")
			logger.Println("Transfer amount:",transferAmount)
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
	helpers.Seperator(1)	
}