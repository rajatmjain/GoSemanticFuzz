package authentication

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func RegisterUser(username string, password string)(){
	account := Account{
		Username: username,
		Password: password,
		Amount: 0,
	}
	fmt.Println("Account created for",account.Username)
	if err:=saveAccount(account);err!=nil{
		fmt.Println("Error:",err)
	}
}
func saveAccount(account Account)(error){
	var accounts []Account
	var initialAmount float64
	float64Schema := goSemanticFuzz.Float64Schema{Minimum: 0,Maximum: 10000,Precision: 2}
	float64Fuzzer := goSemanticFuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
	float64Fuzzer.Fuzz(&initialAmount)
	account.Amount = initialAmount
	if _, err := os.Stat("onlineBanking/storage/accounts.json"); err == nil {
		accountsFile, _ := ioutil.ReadFile("onlineBanking/storage/accounts.json")
		if err := json.Unmarshal([]byte(accountsFile), &accounts);err!=nil{
			return err
		}
		print(accounts)
		accounts = append(accounts, account)
		_, err := json.Marshal(accounts)
		if err!=nil{
			return err
		}
		file, _ := json.MarshalIndent(accounts, "", " ")
		err = ioutil.WriteFile("onlineBanking/storage/accounts.json", file, 0644)
		return err
	 } else {
		accounts = append(accounts, account)
		file, _ := json.MarshalIndent(accounts, "", " ")
		err := ioutil.WriteFile("onlineBanking/storage/accounts.json", file, 0644)
		return err
	}
}
type Account struct{
	Username string `json:"username"`
	Password string `json:"password"`
	Amount float64 `json:"amount"`
}