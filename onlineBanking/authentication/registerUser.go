package authentication

import (
	"fmt"
)

func RegisterUser(username string, password string)(){
	account := make(map[string]string)
	account[username] = password
	fmt.Println("Account created for",username)
	fmt.Println(account)

}