package authentication

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
	"GoSemanticFuzz/onlineBanking/validators"
	"fmt"

	gofuzz "github.com/google/gofuzz"
)
func CreateUserFromGoSemanticFuzz()(error){
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
		fmt.Println("CREATING USER (GoSemanticFuzz)")
		fmt.Println("------------------------------------")
		fmt.Println("Username:",username,"\nPassword:",password)	
		RegisterUser(username,password)
		return err
	}
	return err
}

func CreateUserFromGoFuzz()(error){
	var username string
	usernameUnicodeRange := gofuzz.UnicodeRange{First: '0', Last: 'z'}
	usernameFuzzer := gofuzz.New().Funcs(usernameUnicodeRange.CustomStringFuzzFunc())
	usernameFuzzer.Fuzz(&username)
	var password string
	passwordUnicodeRange := gofuzz.UnicodeRange{First: '0', Last: 'z'}
	passwordFuzzer := gofuzz.New().Funcs(passwordUnicodeRange.CustomStringFuzzFunc())
	passwordFuzzer.Fuzz(&password)
	err := validators.ValidateCredentials(username,password)
	if err==nil{
		fmt.Println("------------------------------------")
		fmt.Println("CREATING USER (GoFuzz)")
		fmt.Println("------------------------------------")
		fmt.Println("Username:",username,"\nPassword:",password)	
		RegisterUser(username,password)
		return err
	}
	return err
}