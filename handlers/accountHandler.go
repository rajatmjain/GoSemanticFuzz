package handlers

import (
	"GoSemanticFuzz/onlineBanking/authentication"
	"GoSemanticFuzz/onlineBanking/misc"
	"fmt"
)

func GoSemanticFuzzAccountGenerator(numberOfTestAccount int){
	misc.Seperator(1)
	fmt.Println("GENERATING",numberOfTestAccount,"TEST ACCOUNT")
	misc.Seperator(1)
	attempts := 0
	i := 0
	for i<numberOfTestAccount{
		if err:=authentication.CreateUserFromGoSemanticFuzz();err==nil{
			i += 1
		}
		attempts += 1
	}
	misc.Seperator(1)
	colorReset := "\033[0m"
    colorRed := "\033[31m"
	fmt.Print(string(colorRed),"Attempts made to create ",numberOfTestAccount," valid accounts: ",attempts,string(colorReset),"\n")
	misc.Seperator(1)
}

func GoFuzzAccountGenerator(numberOfTestAccount int){
	misc.Seperator(1)
	fmt.Println("GENERATING",numberOfTestAccount,"TEST ACCOUNT")
	misc.Seperator(1)
	attempts := 0
	i := 0
	for i<numberOfTestAccount{
		if err:=authentication.CreateUserFromGoFuzz();err==nil{
			i += 1
		}
		attempts += 1
	}
	misc.Seperator(1)
	colorReset := "\033[0m"
    colorRed := "\033[31m"
	fmt.Print(string(colorRed),"Attempts made to create ",numberOfTestAccount," valid accounts: ",attempts,string(colorReset),"\n")
	misc.Seperator(1)
}

func GoSemanticFuzzAccountGeneratorEvaluator(numberOfAttempts int){
	misc.Seperator(1)
	fmt.Println("Number of attempts:",numberOfAttempts)
	misc.Seperator(1)
	valid := 0
	i := 0
	for i<numberOfAttempts{
		if err:=authentication.CreateUserFromGoSemanticFuzz();err==nil{
			valid += 1
		}
		i += 1
	}
	misc.Seperator(1)
	colorReset := "\033[0m"
    colorRed := "\033[31m"
	fmt.Print(string(colorRed),"Attempts made:",numberOfAttempts,", Valid account:",valid,string(colorReset),"\n")
	misc.Seperator(1)
}

func GoFuzzAccountGeneratorEvaluator(numberOfAttempts int){
	misc.Seperator(1)
	fmt.Println("Number of attempts:",numberOfAttempts)
	misc.Seperator(1)
	valid := 0
	i := 0
	for i<numberOfAttempts{
		if err:=authentication.CreateUserFromGoFuzz();err==nil{
			valid += 1
		}
		i += 1
	}
	misc.Seperator(1)
	colorReset := "\033[0m"
    colorRed := "\033[31m"
	fmt.Print(string(colorRed),"Attempts made:",numberOfAttempts,", Valid account:",valid,string(colorReset),"\n")
	misc.Seperator(1)
}