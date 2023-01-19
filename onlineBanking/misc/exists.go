package misc

import (
	"GoSemanticFuzz/onlineBanking/authentication"
	"encoding/json"
	"os"
)

func ValidateAccount(username string)(bool){
	var accounts []authentication.Account
	if _, err := os.Stat("onlineBanking/storage/accounts.json"); err == nil {
		accountsFile, _ := os.ReadFile("onlineBanking/storage/accounts.json")
		if err := json.Unmarshal([]byte(accountsFile), &accounts);err!=nil{
			return false
		}
		for _,acc := range accounts{
			if(acc.Username==username){
				return true
			}
		}
		return false
	}
	return false
}