package validators

import (
	"GoSemanticFuzz/onlineBanking/helpers"
	"errors"
	"fmt"
	"strings"
)

func ValidateCredentials(username string, password string)(error){
	// Assertions
	colorGreen := "\033[32m"
	colorReset := "\033[0m"
	if(username=="" || password==""){
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username and Password empty\n",string(colorReset))
		return errors.New("empty username or password")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username and Password not empty\n",string(colorReset))
	}

	if(strings.ContainsAny(username[0:1],`"!#$%&'()*+,-./):;<=>?@0123456789`)){
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username starts with special character\n",string(colorReset))
		return errors.New("invalid username. Username can not start with an invalid character")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username does not start with invalid character\n",string(colorReset))
	}

	if(len(username)<=2 || len(username)>=16){
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username length invalid\n",string(colorReset))
		return errors.New("invalid username length. Length should be greater than 2 and less than 16")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username length valid\n",string(colorReset))
	}

	if(len(password)<6){
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Password length invalid\n",string(colorReset))
		return errors.New("invalid password length. Length should be greater than 6")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Password length valid\n",string(colorReset))
	}

	if(strings.ContainsAny(username,`"!#$%&'()*+,-./):<=>?@;\?`)){
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username contains invalid characters\n",string(colorReset))
		return errors.New("invalid characters in username")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username contains valid characters\n",string(colorReset))
	}

	if(!strings.ContainsAny(password,`"!#$%&'()*+,-./):<=>?@0123456789;`)){
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Password does not contain required characters\n",string(colorReset))
		return errors.New("missing compulsory characters in password")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Password contains required characters\n",string(colorReset))
	}

	if(username==password){
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username and Password can not be same\n",string(colorReset))
		return errors.New("username and password can not be the same")
	}else{
		helpers.CountAssertions()
		fmt.Print(string(colorGreen),"Username and Password are not same\n",string(colorReset))
	}

	return nil
}