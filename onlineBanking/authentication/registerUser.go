package authentication

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
	"encoding/json"
	"fmt"
	"os"
)

func RegisterUser(username string, password string)(){
	var initialAmount float64
	float64Schema := goSemanticFuzz.Float64Schema{Minimum: 0,Maximum: 10000,Precision: 2}
	float64Fuzzer := goSemanticFuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
	float64Fuzzer.Fuzz(&initialAmount)

	account := Account{
		Username: username,
		AccountDetails: Details{
			Password: password,
			Amount: initialAmount,
		},
	}
	fmt.Println("Account created for",username,"with initial amount $",initialAmount)
	if err:=saveInitialAccount(account);err!=nil{
		fmt.Println("Error:",err)
	}	
}
func saveInitialAccount(account Account)(error){
	var accounts []Account
	if _, err := os.Stat("onlineBanking/storage/accounts.json"); err == nil {
		accountsFile, _ := os.ReadFile("onlineBanking/storage/accounts.json")
		if err := json.Unmarshal([]byte(accountsFile), &accounts);err!=nil{
			return err
		}
		accounts = append(accounts, account)
		_, err := json.Marshal(accounts)
		if err!=nil{
			return err
		}
		file, _ := json.MarshalIndent(accounts, "", " ")
		err = os.WriteFile("onlineBanking/storage/accounts.json", file, 0644)
		return err
	 } else {
		accounts = append(accounts, account)
		file, _ := json.MarshalIndent(accounts, "", " ")
		err := os.WriteFile("onlineBanking/storage/accounts.json", file, 0644)
		return err
	}
}
type Account struct{
	Username string `json:"username"`
	AccountDetails Details `json:"accountDetails"`
}

type Details struct{
	Password string `json:"password"`
	Amount float64 `json:"amount"`
}