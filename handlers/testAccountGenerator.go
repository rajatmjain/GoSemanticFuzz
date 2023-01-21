package handlers

import (
	"GoSemanticFuzz/onlineBanking/authentication"
	"GoSemanticFuzz/onlineBanking/misc"
	"fmt"
)

func TestAccountGenerator(numberOfTestAccount int){
	misc.Seperator(1)
	fmt.Println("GENERATING",numberOfTestAccount,"TEST ACCOUNT")
	misc.Seperator(1)
	attempts := 0
	i := 0
	for i<numberOfTestAccount{
		if err:=authentication.CreateUserFromSemanticFuzz();err==nil{
			i += 1
		}
		attempts += 1
	}
	misc.Seperator(1)
	fmt.Println("Attempts made to create",numberOfTestAccount,"valid account:",attempts)
	misc.Seperator(1)
}