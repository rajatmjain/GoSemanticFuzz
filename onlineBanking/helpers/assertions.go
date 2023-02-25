package helpers

import (
	"os"
	"strconv"
)

func InitAssertionsCount(){
	os.Setenv("assertionCount","0")
}

func CountAssertions(){
	currentAssertionCountString,_ := os.LookupEnv("assertionCount")
	currentAssertionCount,_ := strconv.ParseInt(currentAssertionCountString,10,64)
	currentAssertionCount += 1
	os.Setenv("assertionCount",strconv.Itoa(int(currentAssertionCount)))
}