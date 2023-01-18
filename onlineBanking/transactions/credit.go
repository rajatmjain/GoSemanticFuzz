package transactions

import (
	goSemanticFuzz "GoSemanticFuzz/gofuzz"
	"GoSemanticFuzz/onlineBanking/authentication"
	"fmt"
)

func Credit(account authentication.Account){
	var creditAmount float64
	float64Schema := goSemanticFuzz.Float64Schema{Minimum: 0,Maximum: 10000,Precision: 9}
	float64Fuzzer := goSemanticFuzz.New().Funcs(float64Schema.CustomFloat64FuzzFunc())
	float64Fuzzer.Fuzz(&creditAmount)
	fmt.Println(account.Username,"credits with $",creditAmount) 
}