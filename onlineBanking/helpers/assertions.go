package helpers

import (
	"os"
	"strconv"
)

func InitAssertionsCount(){
	os.Setenv("count","0")
}

func CountAssertions(){
	currentAssertionCountString,_ := os.LookupEnv("count")
	currentAssertionCount,_ := strconv.ParseInt(currentAssertionCountString,10,64)
	currentAssertionCount += 1
	os.Setenv("count",strconv.Itoa(int(currentAssertionCount)))
}