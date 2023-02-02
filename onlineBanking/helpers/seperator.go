package helpers

import (
	"fmt"
	"log"
	"os"
)

func Seperator(count int){
	file, err := os.OpenFile("result.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	logger := log.New(os.Stdout,"",log.LstdFlags)
	logger.SetOutput(file)
	for i:=0;i<count;i++{
		fmt.Println("---------------------------------------------------------------------------")
		logger.Println("---------------------------------------------------------------------------")
	}	
}