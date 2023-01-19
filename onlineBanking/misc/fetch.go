package misc

import (
	"GoSemanticFuzz/onlineBanking/authentication"
	"encoding/json"
	"fmt"
	"os"
)

func FetchAccount(username string)(authentication.Account){
	var accounts []authentication.Account
	if _, err := os.Stat("onlineBanking/storage/accounts.json"); err == nil {
		accountsFile, _ := os.ReadFile("onlineBanking/storage/accounts.json")
		if err := json.Unmarshal([]byte(accountsFile), &accounts);err!=nil{
			fmt.Println("Error unmarshaling")
		}
		for _,acc := range accounts{
			if(acc.Username==username){
				return acc
			}
		}
	}
	return authentication.Account{}
}