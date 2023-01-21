package validators

import (
	"errors"
	"strings"
)

func ValidateCredentials(username string, password string)(error){
	// Assertions
	if(username=="" || password==""){
		return errors.New("empty username or password")
	}
	if(strings.ContainsAny(username[0:1],`"!#$%&'()*+,-./):;<=>?@0123456789`)){
		return errors.New("invalid username. Username can not start with a special character")
	}
	if(len(username)<=2 || len(username)>=16){
		return errors.New("invalid username length. Length should be greater than 2 and less than 16")
	}
	if(len(password)<6){
		return errors.New("invalid password length. Length should be greater than 6")
	}
	if(strings.ContainsAny(username,`"!#$%&'()*+,-./):<=>?@;\?`)){
		return errors.New("invalid characters in username")
	}
	if(!strings.ContainsAny(password,`"!#$%&'()*+,-./):<=>?@0123456789;`)){
		return errors.New("missing compulsory characters in password")
	}
	if(username==password){
		return errors.New("username and password can not be the same")
	}
	
	return nil
}