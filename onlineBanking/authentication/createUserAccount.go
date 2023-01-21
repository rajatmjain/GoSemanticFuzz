package authentication

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
	"GoSemanticFuzz/onlineBanking/validators"
	"fmt"
)
func CreateUserFromSemanticFuzz()(error){
	
	var username string
	usernameUnicodeRange := goSemanticFuzz.UnicodeRange{First: '0', Last: 'z',MinLength: 2,MaxLength: 16}
	usernameFuzzer := goSemanticFuzz.New().Funcs(usernameUnicodeRange.CustomStringFuzzFunc())
	usernameFuzzer.Fuzz(&username)
	var password string
	passwordUnicodeRange := goSemanticFuzz.UnicodeRange{First: '0', Last: 'z',MinLength: 7,MaxLength: 20}
	passwordFuzzer := goSemanticFuzz.New().Funcs(passwordUnicodeRange.CustomStringFuzzFunc())
	passwordFuzzer.Fuzz(&password)
	err := validators.ValidateCredentials(username,password)
	if err==nil{
		fmt.Println("------------------------------------")
		fmt.Println("CREATING USER")
		fmt.Println("------------------------------------")
		fmt.Println("Username:",username,"\nPassword:",password)	
		RegisterUser(username,password)
		return err
	}
	return err
}