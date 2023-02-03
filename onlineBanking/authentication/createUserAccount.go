package authentication

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
	"GoSemanticFuzz/onlineBanking/helpers"
	"GoSemanticFuzz/onlineBanking/validators"
	"fmt"
	"log"
	"os"

	gofuzz "github.com/google/gofuzz"
)
func CreateUserFromGoSemanticFuzz()(error){
	// Logger
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"Debit/ASSERTION:",log.LstdFlags)
	logger.SetOutput(file)
	
	var username string
	usernameUnicodeRange := goSemanticFuzz.UnicodeRange{First: '0', Last: 'z',MinLength: 2,MaxLength: 16}
	usernameFuzzer := goSemanticFuzz.New().Funcs(usernameUnicodeRange.CustomStringFuzzFunc())
	usernameFuzzer.Fuzz(&username)
	var password string
	passwordUnicodeRange := goSemanticFuzz.UnicodeRange{First: '0', Last: 'z',MinLength: 7,MaxLength: 20}
	passwordFuzzer := goSemanticFuzz.New().Funcs(passwordUnicodeRange.CustomStringFuzzFunc())
	passwordFuzzer.Fuzz(&password)
	err = validators.ValidateCredentials(username,password)
	if err==nil{
		helpers.Seperator(1)
		fmt.Println("CREATING USER (GoSemanticFuzz)")
		helpers.Seperator(1)
		fmt.Println("Username:",username,"\nPassword:",password)	
		RegisterUser(username,password)
		// Assertion
		helpers.CountAssertions()
		colorGreen := "\033[32m"
		colorReset := "\033[0m"
		fmt.Print(string(colorGreen),"Valid account created","\n",string(colorReset))
		logger.Println("Valid account created")
		return err
	}
	// Assertion
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	helpers.Seperator(1)
	fmt.Print(string(colorGreen),"Invalid account created","\n",string(colorReset))
	logger.Println("Invalid account created")
	helpers.Seperator(1)
	return err
}

func CreateUserFromGoFuzz()(error){
	// Logger
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"Debit/ASSERTION:",log.LstdFlags)
	logger.SetOutput(file)

	var username string
	usernameUnicodeRange := gofuzz.UnicodeRange{First: '0', Last: 'z'}
	usernameFuzzer := gofuzz.New().Funcs(usernameUnicodeRange.CustomStringFuzzFunc())
	usernameFuzzer.Fuzz(&username)
	var password string
	passwordUnicodeRange := gofuzz.UnicodeRange{First: '0', Last: 'z'}
	passwordFuzzer := gofuzz.New().Funcs(passwordUnicodeRange.CustomStringFuzzFunc())
	passwordFuzzer.Fuzz(&password)
	err = validators.ValidateCredentials(username,password)
	if err==nil{
		helpers.Seperator(1)
		fmt.Println("CREATING USER (GoFuzz)")
		helpers.Seperator(1)
		fmt.Println("Username:",username,"\nPassword:",password)	
		RegisterUser(username,password)
		// Assertion
		helpers.CountAssertions()
		colorGreen := "\033[32m"
		colorReset := "\033[0m"
		fmt.Print(string(colorGreen),"Valid account created","\n",string(colorReset))
		logger.Println("Valid account created")
		return err
	}
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	helpers.Seperator(1)
	fmt.Print(string(colorGreen),"Invalid account created","\n",string(colorReset))
	logger.Println("Invalid account created")
	helpers.Seperator(1)
	return err
}