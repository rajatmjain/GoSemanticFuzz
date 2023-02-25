package helpers

import (
	"os"
	"strconv"
)

func InitBranchesCount(){
	os.Setenv("branchCount","0")
}

func CountBranches(){
	currentBranchCountString,_ := os.LookupEnv("branchCount")
	currentBranchCount,_ := strconv.ParseInt(currentBranchCountString,10,64)
	currentBranchCount += 1
	os.Setenv("branchCount",strconv.Itoa(int(currentBranchCount)))
}