package misc

import (
	"GoSemanticFuzz/onlineBanking/authentication"
	"encoding/json"
	"errors"
	"os"
)

func SaveAccount(account authentication.Account)(error){
	var accounts []authentication.Account
	if _, err := os.Stat("onlineBanking/storage/accounts.json"); err == nil {
		accountsFile, _ := os.ReadFile("onlineBanking/storage/accounts.json")
		if err := json.Unmarshal([]byte(accountsFile), &accounts);err!=nil{
			return err
		}
		for index,acc := range accounts{
			if(acc.Username==account.Username){
				accounts = append(accounts[:index],accounts[index+1:]...)
				accounts = append(accounts, account)
				_, err := json.Marshal(accounts)
				if err!=nil{
					return err
				}
				file, _ := json.MarshalIndent(accounts, "", " ")
				err = os.WriteFile("onlineBanking/storage/accounts.json", file, 0644)
				return err
			}
		}
		return errors.New("account does not exist")
	}
	// 	accounts = append(accounts, account)
	// 	_, err := json.Marshal(accounts)
	// 	if err!=nil{
	// 		return err
	// 	}
	// 	file, _ := json.MarshalIndent(accounts, "", " ")
	// 	err = os.WriteFile("onlineBanking/storage/accounts.json", file, 0644)
	// 	return err
	//  } else {
	// 	accounts = append(accounts, account)
	// 	file, _ := json.MarshalIndent(accounts, "", " ")
	// 	err := os.WriteFile("onlineBanking/storage/accounts.json", file, 0644)
	// 	return err
	// }
	return nil
}