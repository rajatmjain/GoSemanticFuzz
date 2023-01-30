package handlers

import (
	"GoSemanticFuzz/onlineBanking/authentication"
	"GoSemanticFuzz/onlineBanking/helpers"
	"fmt"
)

func GoSemanticFuzzAccountGenerator(numberOfTestAccount int){
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
	helpers.Seperator(1)
}

func GoFuzzAccountGenerator(numberOfTestAccount int){
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
	helpers.Seperator(1)
}

func GoSemanticFuzzAccountGeneratorEvaluator(numberOfAttempts int){
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
	helpers.Seperator(1)
}

func GoFuzzAccountGeneratorEvaluator(numberOfAttempts int){
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
	helpers.Seperator(1)
}