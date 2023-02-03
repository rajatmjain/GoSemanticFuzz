package validators

import (
	"GoSemanticFuzz/onlineBanking/helpers"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func ValidateCredentials(username string, password string)(error){
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"CV/ASSERTION:",log.LstdFlags)
	logger.SetOutput(file)
	// Assertions
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	if(username=="" || password==""){
		fmt.Print(string(colorGreen),"Username and Password empty\n",string(colorReset))
		logger.Print("Username and Password empty\n")
		return errors.New("empty username or password")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username and Password not empty\n",string(colorReset))
		logger.Print("Username and Password not empty\n")
	}

	if(strings.ContainsAny(username[0:1],`"!#$%&'()*+,-./):;<=>?@0123456789`)){
		fmt.Print(string(colorGreen),"Username starts with special character\n",string(colorReset))
		logger.Print("Username starts with special character\n")
		return errors.New("invalid username. Username can not start with an invalid character")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username does not start with invalid character\n",string(colorReset))
		logger.Print("Username does not start with invalid character\n")
	}

	if(len(username)<=2 || len(username)>=16){
		fmt.Print(string(colorGreen),"Username length invalid\n",string(colorReset))
		logger.Print("Username length invalid\n")
		return errors.New("invalid username length. Length should be greater than 2 and less than 16")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username length valid\n",string(colorReset))
		logger.Print("Username length valid\n")
	}

	if(len(password)<6){
		fmt.Print(string(colorGreen),"Password length invalid\n",string(colorReset))
		logger.Print("Password length invalid\n")
		return errors.New("invalid password length. Length should be greater than 6")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Password length valid\n",string(colorReset))
		logger.Print("Password length valid\n")
	}

	if(strings.ContainsAny(username,`"!#$%&'()*+,-./):<=>?@;\?`)){
		fmt.Print(string(colorGreen),"Username contains invalid characters\n",string(colorReset))
		logger.Print("Username contains invalid characters\n")
		return errors.New("invalid characters in username")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username contains valid characters\n",string(colorReset))
		logger.Print("Username contains valid characters\n")
	}

	if(!strings.ContainsAny(password,`"!#$%&'()*+,-./):<=>?@0123456789;`)){
		fmt.Print(string(colorGreen),"Password does not contain required characters\n",string(colorReset))
		logger.Print("Password does not contain required characters\n")
		return errors.New("missing compulsory characters in password")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Password contains required characters\n",string(colorReset))
		logger.Print("Password contains required characters\n")
	}

	if(username==password){
		fmt.Print(string(colorGreen),"Username and Password can not be same\n",string(colorReset))
		logger.Print("Username and Password can not be same\n")
		return errors.New("username and password can not be the same")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username and Password are not same\n",string(colorReset))
		logger.Print("Username and Password are not same\n")
	}
	return nil
}