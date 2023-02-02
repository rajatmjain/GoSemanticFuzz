package handlers

import (
	"GoSemanticFuzz/onlineBanking/authentication"
	"GoSemanticFuzz/onlineBanking/helpers"
	"fmt"
	"log"
	"os"
)

func GoSemanticFuzzAccountGenerator(numberOfTestAccount int){
	// LOGGER //
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"GSFTestRoutine:",log.LstdFlags)
	logger.SetOutput(file)

	helpers.Seperator(1)
	fmt.Println("GENERATING",numberOfTestAccount,"TEST ACCOUNT")
	helpers.Seperator(1)
	attempts := 0
	i := 0
	for i<numberOfTestAccount{
		if err:=authentication.CreateUserFromGoSemanticFuzz();err==nil{
			i += 1
		}
		attempts += 1
	}
	helpers.Seperator(1)
	colorReset := "\033[0m"
    colorRed := "\033[31m"
	fmt.Print(string(colorRed),"Attempts made to create ",numberOfTestAccount," valid accounts: ",attempts,string(colorReset),"\n")
	logger.Println("Attempts made to create ",numberOfTestAccount," valid accounts: ",attempts)
	helpers.Seperator(1)
}

func GoFuzzAccountGenerator(numberOfTestAccount int){
	// LOGGER //
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"GFTestRoutine:",log.LstdFlags)
	logger.SetOutput(file)
	helpers.Seperator(1)
	fmt.Println("GENERATING",numberOfTestAccount,"TEST ACCOUNT")
	helpers.Seperator(1)
	attempts := 0
	i := 0
	for i<numberOfTestAccount{
		if err:=authentication.CreateUserFromGoFuzz();err==nil{
			i += 1
		}
		attempts += 1
	}
	helpers.Seperator(1)
	colorReset := "\033[0m"
    colorRed := "\033[31m"
	fmt.Print(string(colorRed),"Attempts made to create ",numberOfTestAccount," valid accounts: ",attempts,string(colorReset),"\n")
	logger.Println("Attempts made to create ",numberOfTestAccount," valid accounts: ",attempts)
	helpers.Seperator(1)
}

func GoSemanticFuzzAccountGeneratorEvaluator(numberOfAttempts int){
	// LOGGER //
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"GSFTestRoutine:",log.LstdFlags)
	logger.SetOutput(file)

	helpers.Seperator(1)
	fmt.Println("Number of attempts:",numberOfAttempts)
	helpers.Seperator(1)
	valid := 0
	i := 0
	for i<numberOfAttempts{
		if err:=authentication.CreateUserFromGoSemanticFuzz();err==nil{
			valid += 1
		}
		i += 1
	}
	helpers.Seperator(1)
	colorReset := "\033[0m"
    colorRed := "\033[31m"
	fmt.Print(string(colorRed),"Attempts made:",numberOfAttempts,", Valid account:",valid,string(colorReset),"\n")
	logger.Println("Attempts made:",numberOfAttempts,", Valid account:",valid)
	helpers.Seperator(1)
}

func GoFuzzAccountGeneratorEvaluator(numberOfAttempts int){
	// LOGGER //
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"GFTestRoutine:",log.LstdFlags)
	logger.SetOutput(file)

	helpers.Seperator(1)
	fmt.Println("Number of attempts:",numberOfAttempts)
	helpers.Seperator(1)
	valid := 0
	i := 0
	for i<numberOfAttempts{
		if err:=authentication.CreateUserFromGoFuzz();err==nil{
			valid += 1
		}
		i += 1
	}
	helpers.Seperator(1)
	colorReset := "\033[0m"
    colorRed := "\033[31m"
	fmt.Print(string(colorRed),"Attempts made:",numberOfAttempts,", Valid account:",valid,string(colorReset),"\n")
	logger.Println("Attempts made:",numberOfAttempts,", Valid account:",valid)
	helpers.Seperator(1)
}